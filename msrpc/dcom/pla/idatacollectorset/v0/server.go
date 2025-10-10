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
		op := &xxx_GetDataCollectorsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDataCollectorsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDataCollectors(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Duration
		op := &xxx_GetDurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDuration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Duration
		op := &xxx_SetDurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDuration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Description
		op := &xxx_GetDescriptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDescriptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDescription(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Description
		op := &xxx_SetDescriptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDescriptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDescription(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // DescriptionUnresolved
		op := &xxx_GetDescriptionUnresolvedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDescriptionUnresolvedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDescriptionUnresolved(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // DisplayName
		op := &xxx_GetDisplayNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDisplayNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDisplayName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // DisplayName
		op := &xxx_SetDisplayNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDisplayNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDisplayName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // DisplayNameUnresolved
		op := &xxx_GetDisplayNameUnresolvedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDisplayNameUnresolvedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDisplayNameUnresolved(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // Keywords
		op := &xxx_GetKeywordsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetKeywordsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetKeywords(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // Keywords
		op := &xxx_SetKeywordsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetKeywordsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetKeywords(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // LatestOutputLocation
		op := &xxx_GetLatestOutputLocationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLatestOutputLocationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLatestOutputLocation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // LatestOutputLocation
		op := &xxx_SetLatestOutputLocationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLatestOutputLocationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLatestOutputLocation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // Name
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // OutputLocation
		op := &xxx_GetOutputLocationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOutputLocationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOutputLocation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // RootPath
		op := &xxx_GetRootPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRootPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRootPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // RootPath
		op := &xxx_SetRootPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetRootPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetRootPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // Segment
		op := &xxx_GetSegmentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSegmentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSegment(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // Segment
		op := &xxx_SetSegmentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSegmentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSegment(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // SegmentMaxDuration
		op := &xxx_GetSegmentMaxDurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSegmentMaxDurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSegmentMaxDuration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // SegmentMaxDuration
		op := &xxx_SetSegmentMaxDurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSegmentMaxDurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSegmentMaxDuration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // SegmentMaxSize
		op := &xxx_GetSegmentMaxSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSegmentMaxSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSegmentMaxSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // SegmentMaxSize
		op := &xxx_SetSegmentMaxSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSegmentMaxSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSegmentMaxSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // SerialNumber
		op := &xxx_GetSerialNumberOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSerialNumberRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSerialNumber(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // SerialNumber
		op := &xxx_SetSerialNumberOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSerialNumberRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSerialNumber(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // Server
		op := &xxx_GetServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // Status
		op := &xxx_GetStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // Subdirectory
		op := &xxx_GetSubdirectoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSubdirectoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSubdirectory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // Subdirectory
		op := &xxx_SetSubdirectoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSubdirectoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSubdirectory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // SubdirectoryFormat
		op := &xxx_GetSubdirectoryFormatOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSubdirectoryFormatRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSubdirectoryFormat(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // SubdirectoryFormat
		op := &xxx_SetSubdirectoryFormatOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSubdirectoryFormatRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSubdirectoryFormat(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // SubdirectoryFormatPattern
		op := &xxx_GetSubdirectoryFormatPatternOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSubdirectoryFormatPatternRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSubdirectoryFormatPattern(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // SubdirectoryFormatPattern
		op := &xxx_SetSubdirectoryFormatPatternOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSubdirectoryFormatPatternRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSubdirectoryFormatPattern(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // Task
		op := &xxx_GetTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // Task
		op := &xxx_SetTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // TaskRunAsSelf
		op := &xxx_GetTaskRunAsSelfOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTaskRunAsSelfRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTaskRunAsSelf(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // TaskRunAsSelf
		op := &xxx_SetTaskRunAsSelfOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTaskRunAsSelfRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTaskRunAsSelf(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // TaskArguments
		op := &xxx_GetTaskArgumentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTaskArgumentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTaskArguments(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // TaskArguments
		op := &xxx_SetTaskArgumentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTaskArgumentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTaskArguments(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // TaskUserTextArguments
		op := &xxx_GetTaskUserTextArgumentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTaskUserTextArgumentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTaskUserTextArguments(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // TaskUserTextArguments
		op := &xxx_SetTaskUserTextArgumentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTaskUserTextArgumentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTaskUserTextArguments(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // Schedules
		op := &xxx_GetSchedulesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSchedulesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSchedules(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // SchedulesEnabled
		op := &xxx_GetSchedulesEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSchedulesEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSchedulesEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 50: // SchedulesEnabled
		op := &xxx_SetSchedulesEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSchedulesEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSchedulesEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 51: // UserAccount
		op := &xxx_GetUserAccountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUserAccountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUserAccount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 52: // Xml
		op := &xxx_GetXMLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetXMLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetXML(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 53: // Security
		op := &xxx_GetSecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 54: // Security
		op := &xxx_SetSecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 55: // StopOnCompletion
		op := &xxx_GetStopOnCompletionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetStopOnCompletionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetStopOnCompletion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 56: // StopOnCompletion
		op := &xxx_SetStopOnCompletionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetStopOnCompletionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetStopOnCompletion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 57: // DataManager
		op := &xxx_GetDataManagerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDataManagerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDataManager(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 58: // SetCredentials
		op := &xxx_SetCredentialsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCredentialsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCredentials(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 59: // Query
		op := &xxx_QueryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Query(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 60: // Commit
		op := &xxx_CommitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CommitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Commit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 61: // Delete
		op := &xxx_DeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Delete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 62: // Start
		op := &xxx_StartOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Start(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 63: // Stop
		op := &xxx_StopOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StopRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Stop(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 64: // SetXml
		op := &xxx_SetXMLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetXMLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetXML(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 65: // SetValue
		op := &xxx_SetValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 66: // GetValue
		op := &xxx_GetValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IDataCollectorSet
type UnimplementedDataCollectorSetServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedDataCollectorSetServer) GetDataCollectors(context.Context, *GetDataCollectorsRequest) (*GetDataCollectorsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetDuration(context.Context, *GetDurationRequest) (*GetDurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetDuration(context.Context, *SetDurationRequest) (*SetDurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetDescription(context.Context, *SetDescriptionRequest) (*SetDescriptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetDescriptionUnresolved(context.Context, *GetDescriptionUnresolvedRequest) (*GetDescriptionUnresolvedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetDisplayName(context.Context, *GetDisplayNameRequest) (*GetDisplayNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetDisplayName(context.Context, *SetDisplayNameRequest) (*SetDisplayNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetDisplayNameUnresolved(context.Context, *GetDisplayNameUnresolvedRequest) (*GetDisplayNameUnresolvedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetKeywords(context.Context, *GetKeywordsRequest) (*GetKeywordsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetKeywords(context.Context, *SetKeywordsRequest) (*SetKeywordsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetLatestOutputLocation(context.Context, *GetLatestOutputLocationRequest) (*GetLatestOutputLocationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetLatestOutputLocation(context.Context, *SetLatestOutputLocationRequest) (*SetLatestOutputLocationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetOutputLocation(context.Context, *GetOutputLocationRequest) (*GetOutputLocationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetRootPath(context.Context, *GetRootPathRequest) (*GetRootPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetRootPath(context.Context, *SetRootPathRequest) (*SetRootPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetSegment(context.Context, *GetSegmentRequest) (*GetSegmentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetSegment(context.Context, *SetSegmentRequest) (*SetSegmentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetSegmentMaxDuration(context.Context, *GetSegmentMaxDurationRequest) (*GetSegmentMaxDurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetSegmentMaxDuration(context.Context, *SetSegmentMaxDurationRequest) (*SetSegmentMaxDurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetSegmentMaxSize(context.Context, *GetSegmentMaxSizeRequest) (*GetSegmentMaxSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetSegmentMaxSize(context.Context, *SetSegmentMaxSizeRequest) (*SetSegmentMaxSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetSerialNumber(context.Context, *GetSerialNumberRequest) (*GetSerialNumberResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetSerialNumber(context.Context, *SetSerialNumberRequest) (*SetSerialNumberResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetServer(context.Context, *GetServerRequest) (*GetServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetSubdirectory(context.Context, *GetSubdirectoryRequest) (*GetSubdirectoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetSubdirectory(context.Context, *SetSubdirectoryRequest) (*SetSubdirectoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetSubdirectoryFormat(context.Context, *GetSubdirectoryFormatRequest) (*GetSubdirectoryFormatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetSubdirectoryFormat(context.Context, *SetSubdirectoryFormatRequest) (*SetSubdirectoryFormatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetSubdirectoryFormatPattern(context.Context, *GetSubdirectoryFormatPatternRequest) (*GetSubdirectoryFormatPatternResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetSubdirectoryFormatPattern(context.Context, *SetSubdirectoryFormatPatternRequest) (*SetSubdirectoryFormatPatternResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetTask(context.Context, *SetTaskRequest) (*SetTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetTaskRunAsSelf(context.Context, *GetTaskRunAsSelfRequest) (*GetTaskRunAsSelfResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetTaskRunAsSelf(context.Context, *SetTaskRunAsSelfRequest) (*SetTaskRunAsSelfResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetTaskArguments(context.Context, *GetTaskArgumentsRequest) (*GetTaskArgumentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetTaskArguments(context.Context, *SetTaskArgumentsRequest) (*SetTaskArgumentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetTaskUserTextArguments(context.Context, *GetTaskUserTextArgumentsRequest) (*GetTaskUserTextArgumentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetTaskUserTextArguments(context.Context, *SetTaskUserTextArgumentsRequest) (*SetTaskUserTextArgumentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetSchedules(context.Context, *GetSchedulesRequest) (*GetSchedulesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetSchedulesEnabled(context.Context, *GetSchedulesEnabledRequest) (*GetSchedulesEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetSchedulesEnabled(context.Context, *SetSchedulesEnabledRequest) (*SetSchedulesEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetUserAccount(context.Context, *GetUserAccountRequest) (*GetUserAccountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetXML(context.Context, *GetXMLRequest) (*GetXMLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetSecurity(context.Context, *SetSecurityRequest) (*SetSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetStopOnCompletion(context.Context, *GetStopOnCompletionRequest) (*GetStopOnCompletionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetStopOnCompletion(context.Context, *SetStopOnCompletionRequest) (*SetStopOnCompletionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetDataManager(context.Context, *GetDataManagerRequest) (*GetDataManagerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetCredentials(context.Context, *SetCredentialsRequest) (*SetCredentialsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetXML(context.Context, *SetXMLRequest) (*SetXMLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetServer) GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DataCollectorSetServer = (*UnimplementedDataCollectorSetServer)(nil)
