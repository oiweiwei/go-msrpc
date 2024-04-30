package idatacollectorset

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

// IDataCollectorSet server interface.
type DataCollectorSetServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The DataCollectors (Get) method retrieves the DataCollectors property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetDataCollectors(context.Context, *GetDataCollectorsRequest) (*GetDataCollectorsResponse, error)

	// Duration operation.
	GetDuration(context.Context, *GetDurationRequest) (*GetDurationResponse, error)

	// Duration operation.
	SetDuration(context.Context, *SetDurationRequest) (*SetDurationResponse, error)

	// Description operation.
	GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error)

	// Description operation.
	SetDescription(context.Context, *SetDescriptionRequest) (*SetDescriptionResponse, error)

	// The DescriptionUnresolved (Get) method retrieves the DescriptionUnresolved property.
	//
	// This method MUST return the description as originally set by the IDataCollectorSet::Description
	// method, as specified in section 3.2.4.1.5.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetDescriptionUnresolved(context.Context, *GetDescriptionUnresolvedRequest) (*GetDescriptionUnresolvedResponse, error)

	// DisplayName operation.
	GetDisplayName(context.Context, *GetDisplayNameRequest) (*GetDisplayNameResponse, error)

	// DisplayName operation.
	SetDisplayName(context.Context, *SetDisplayNameRequest) (*SetDisplayNameResponse, error)

	// The DisplayNameUnresolved (Get) method retrieves the DisplayNameUnresolved property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetDisplayNameUnresolved(context.Context, *GetDisplayNameUnresolvedRequest) (*GetDisplayNameUnresolvedResponse, error)

	// Keywords operation.
	GetKeywords(context.Context, *GetKeywordsRequest) (*GetKeywordsResponse, error)

	// Keywords operation.
	SetKeywords(context.Context, *SetKeywordsRequest) (*SetKeywordsResponse, error)

	// LatestOutputLocation operation.
	GetLatestOutputLocation(context.Context, *GetLatestOutputLocationRequest) (*GetLatestOutputLocationResponse, error)

	// LatestOutputLocation operation.
	SetLatestOutputLocation(context.Context, *SetLatestOutputLocationRequest) (*SetLatestOutputLocationResponse, error)

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// OutputLocation operation.
	GetOutputLocation(context.Context, *GetOutputLocationRequest) (*GetOutputLocationResponse, error)

	// RootPath operation.
	GetRootPath(context.Context, *GetRootPathRequest) (*GetRootPathResponse, error)

	// RootPath operation.
	SetRootPath(context.Context, *SetRootPathRequest) (*SetRootPathResponse, error)

	// Segment operation.
	GetSegment(context.Context, *GetSegmentRequest) (*GetSegmentResponse, error)

	// Segment operation.
	SetSegment(context.Context, *SetSegmentRequest) (*SetSegmentResponse, error)

	// SegmentMaxDuration operation.
	GetSegmentMaxDuration(context.Context, *GetSegmentMaxDurationRequest) (*GetSegmentMaxDurationResponse, error)

	// SegmentMaxDuration operation.
	SetSegmentMaxDuration(context.Context, *SetSegmentMaxDurationRequest) (*SetSegmentMaxDurationResponse, error)

	// SegmentMaxSize operation.
	GetSegmentMaxSize(context.Context, *GetSegmentMaxSizeRequest) (*GetSegmentMaxSizeResponse, error)

	// SegmentMaxSize operation.
	SetSegmentMaxSize(context.Context, *SetSegmentMaxSizeRequest) (*SetSegmentMaxSizeResponse, error)

	// SerialNumber operation.
	GetSerialNumber(context.Context, *GetSerialNumberRequest) (*GetSerialNumberResponse, error)

	// SerialNumber operation.
	SetSerialNumber(context.Context, *SetSerialNumberRequest) (*SetSerialNumberResponse, error)

	// The Server (Get) method retrieves the Server property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetServer(context.Context, *GetServerRequest) (*GetServerResponse, error)

	// The Status (Get) method retrieves the Status property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)

	// Subdirectory operation.
	GetSubdirectory(context.Context, *GetSubdirectoryRequest) (*GetSubdirectoryResponse, error)

	// Subdirectory operation.
	SetSubdirectory(context.Context, *SetSubdirectoryRequest) (*SetSubdirectoryResponse, error)

	// SubdirectoryFormat operation.
	GetSubdirectoryFormat(context.Context, *GetSubdirectoryFormatRequest) (*GetSubdirectoryFormatResponse, error)

	// SubdirectoryFormat operation.
	SetSubdirectoryFormat(context.Context, *SetSubdirectoryFormatRequest) (*SetSubdirectoryFormatResponse, error)

	// SubdirectoryFormatPattern operation.
	GetSubdirectoryFormatPattern(context.Context, *GetSubdirectoryFormatPatternRequest) (*GetSubdirectoryFormatPatternResponse, error)

	// SubdirectoryFormatPattern operation.
	SetSubdirectoryFormatPattern(context.Context, *SetSubdirectoryFormatPatternRequest) (*SetSubdirectoryFormatPatternResponse, error)

	// Task operation.
	GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error)

	// Task operation.
	SetTask(context.Context, *SetTaskRequest) (*SetTaskResponse, error)

	// TaskRunAsSelf operation.
	GetTaskRunAsSelf(context.Context, *GetTaskRunAsSelfRequest) (*GetTaskRunAsSelfResponse, error)

	// TaskRunAsSelf operation.
	SetTaskRunAsSelf(context.Context, *SetTaskRunAsSelfRequest) (*SetTaskRunAsSelfResponse, error)

	// TaskArguments operation.
	GetTaskArguments(context.Context, *GetTaskArgumentsRequest) (*GetTaskArgumentsResponse, error)

	// TaskArguments operation.
	SetTaskArguments(context.Context, *SetTaskArgumentsRequest) (*SetTaskArgumentsResponse, error)

	// TaskUserTextArguments operation.
	GetTaskUserTextArguments(context.Context, *GetTaskUserTextArgumentsRequest) (*GetTaskUserTextArgumentsResponse, error)

	// TaskUserTextArguments operation.
	SetTaskUserTextArguments(context.Context, *SetTaskUserTextArgumentsRequest) (*SetTaskUserTextArgumentsResponse, error)

	// The Schedules (Get) method retrieves the Schedules property, as specified in the
	// property table in section 3.2.4.1.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetSchedules(context.Context, *GetSchedulesRequest) (*GetSchedulesResponse, error)

	// SchedulesEnabled operation.
	GetSchedulesEnabled(context.Context, *GetSchedulesEnabledRequest) (*GetSchedulesEnabledResponse, error)

	// SchedulesEnabled operation.
	SetSchedulesEnabled(context.Context, *SetSchedulesEnabledRequest) (*SetSchedulesEnabledResponse, error)

	// The UserAccount (Get) method retrieves the UserAccount property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetUserAccount(context.Context, *GetUserAccountRequest) (*GetUserAccountResponse, error)

	// Xml operation.
	GetXML(context.Context, *GetXMLRequest) (*GetXMLResponse, error)

	// Security operation.
	GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error)

	// Security operation.
	SetSecurity(context.Context, *SetSecurityRequest) (*SetSecurityResponse, error)

	// StopOnCompletion operation.
	GetStopOnCompletion(context.Context, *GetStopOnCompletionRequest) (*GetStopOnCompletionResponse, error)

	// StopOnCompletion operation.
	SetStopOnCompletion(context.Context, *SetStopOnCompletionRequest) (*SetStopOnCompletionResponse, error)

	// The DataManager (Get) method retrieves the DataManager property, as specified in
	// the property table in section 3.2.4.1.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetDataManager(context.Context, *GetDataManagerRequest) (*GetDataManagerResponse, error)

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
	SetCredentials(context.Context, *SetCredentialsRequest) (*SetCredentialsResponse, error)

	// Query operation.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)

	// The Commit method updates, validates, or saves a data collector set, or flushes the
	// event trace data collectors of a data collector set.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)

	// The Delete method removes the data collector set from storage if it is not running.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)

	// The Start method manually starts the data collector set.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	Start(context.Context, *StartRequest) (*StartResponse, error)

	// The Stop method manually stops the data collector set.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	Stop(context.Context, *StopRequest) (*StopResponse, error)

	// SetXml operation.
	SetXML(context.Context, *SetXMLRequest) (*SetXMLResponse, error)

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
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)

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
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)
}

func RegisterDataCollectorSetServer(conn dcerpc.Conn, o DataCollectorSetServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDataCollectorSetServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DataCollectorSetSyntaxV0_0))...)
}

func NewDataCollectorSetServerHandle(o DataCollectorSetServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DataCollectorSetServerHandle(ctx, o, opNum, r)
	}
}

func DataCollectorSetServerHandle(ctx context.Context, o DataCollectorSetServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // DataCollectors
		in := &GetDataCollectorsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDataCollectors(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Duration
		in := &GetDurationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDuration(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Duration
		in := &SetDurationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDuration(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Description
		in := &GetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Description
		in := &SetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // DescriptionUnresolved
		in := &GetDescriptionUnresolvedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDescriptionUnresolved(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // DisplayName
		in := &GetDisplayNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDisplayName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // DisplayName
		in := &SetDisplayNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDisplayName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // DisplayNameUnresolved
		in := &GetDisplayNameUnresolvedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDisplayNameUnresolved(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // Keywords
		in := &GetKeywordsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetKeywords(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // Keywords
		in := &SetKeywordsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetKeywords(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // LatestOutputLocation
		in := &GetLatestOutputLocationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLatestOutputLocation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // LatestOutputLocation
		in := &SetLatestOutputLocationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetLatestOutputLocation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // Name
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // OutputLocation
		in := &GetOutputLocationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOutputLocation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // RootPath
		in := &GetRootPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRootPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // RootPath
		in := &SetRootPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetRootPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // Segment
		in := &GetSegmentRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSegment(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // Segment
		in := &SetSegmentRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSegment(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // SegmentMaxDuration
		in := &GetSegmentMaxDurationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSegmentMaxDuration(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // SegmentMaxDuration
		in := &SetSegmentMaxDurationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSegmentMaxDuration(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // SegmentMaxSize
		in := &GetSegmentMaxSizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSegmentMaxSize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // SegmentMaxSize
		in := &SetSegmentMaxSizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSegmentMaxSize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // SerialNumber
		in := &GetSerialNumberRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSerialNumber(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // SerialNumber
		in := &SetSerialNumberRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSerialNumber(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // Server
		in := &GetServerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetServer(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // Status
		in := &GetStatusRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetStatus(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // Subdirectory
		in := &GetSubdirectoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubdirectory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // Subdirectory
		in := &SetSubdirectoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSubdirectory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // SubdirectoryFormat
		in := &GetSubdirectoryFormatRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubdirectoryFormat(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // SubdirectoryFormat
		in := &SetSubdirectoryFormatRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSubdirectoryFormat(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // SubdirectoryFormatPattern
		in := &GetSubdirectoryFormatPatternRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubdirectoryFormatPattern(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // SubdirectoryFormatPattern
		in := &SetSubdirectoryFormatPatternRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSubdirectoryFormatPattern(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // Task
		in := &GetTaskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTask(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // Task
		in := &SetTaskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetTask(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // TaskRunAsSelf
		in := &GetTaskRunAsSelfRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTaskRunAsSelf(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // TaskRunAsSelf
		in := &SetTaskRunAsSelfRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetTaskRunAsSelf(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // TaskArguments
		in := &GetTaskArgumentsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTaskArguments(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // TaskArguments
		in := &SetTaskArgumentsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetTaskArguments(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // TaskUserTextArguments
		in := &GetTaskUserTextArgumentsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTaskUserTextArguments(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // TaskUserTextArguments
		in := &SetTaskUserTextArgumentsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetTaskUserTextArguments(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 48: // Schedules
		in := &GetSchedulesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSchedules(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // SchedulesEnabled
		in := &GetSchedulesEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSchedulesEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 50: // SchedulesEnabled
		in := &SetSchedulesEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSchedulesEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 51: // UserAccount
		in := &GetUserAccountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetUserAccount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 52: // Xml
		in := &GetXMLRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetXML(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 53: // Security
		in := &GetSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 54: // Security
		in := &SetSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 55: // StopOnCompletion
		in := &GetStopOnCompletionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetStopOnCompletion(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 56: // StopOnCompletion
		in := &SetStopOnCompletionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetStopOnCompletion(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 57: // DataManager
		in := &GetDataManagerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDataManager(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 58: // SetCredentials
		in := &SetCredentialsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetCredentials(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 59: // Query
		in := &QueryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Query(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 60: // Commit
		in := &CommitRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Commit(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 61: // Delete
		in := &DeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Delete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 62: // Start
		in := &StartRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Start(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 63: // Stop
		in := &StopRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Stop(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 64: // SetXml
		in := &SetXMLRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetXML(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 65: // SetValue
		in := &SetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 66: // GetValue
		in := &GetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
