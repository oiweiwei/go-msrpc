package ifsrmfilemanagementjob

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

// IFsrmFileManagementJob server interface.
type FileManagementJobServer interface {

	// IFsrmObject base class.
	ifsrmobject.ObjectServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest) (*SetNameResponse, error)

	// NamespaceRoots operation.
	GetNamespaceRoots(context.Context, *GetNamespaceRootsRequest) (*GetNamespaceRootsResponse, error)

	// NamespaceRoots operation.
	SetNamespaceRoots(context.Context, *SetNamespaceRootsRequest) (*SetNamespaceRootsResponse, error)

	// Enabled operation.
	GetEnabled(context.Context, *GetEnabledRequest) (*GetEnabledResponse, error)

	// Enabled operation.
	SetEnabled(context.Context, *SetEnabledRequest) (*SetEnabledResponse, error)

	// OperationType operation.
	GetOperationType(context.Context, *GetOperationTypeRequest) (*GetOperationTypeResponse, error)

	// OperationType operation.
	SetOperationType(context.Context, *SetOperationTypeRequest) (*SetOperationTypeResponse, error)

	// ExpirationDirectory operation.
	GetExpirationDirectory(context.Context, *GetExpirationDirectoryRequest) (*GetExpirationDirectoryResponse, error)

	// ExpirationDirectory operation.
	SetExpirationDirectory(context.Context, *SetExpirationDirectoryRequest) (*SetExpirationDirectoryResponse, error)

	// The CustomAction (get) method retrieves the custom action of the file management
	// job and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-------------------------------+
	//	|         RETURN          |                               |
	//	|       VALUE/CODE        |          DESCRIPTION          |
	//	|                         |                               |
	//	+-------------------------+-------------------------------+
	//	+-------------------------+-------------------------------+
	//	| 0x80070057 E_INVALIDARG | The action parameter is NULL. |
	//	+-------------------------+-------------------------------+
	GetCustomAction(context.Context, *GetCustomActionRequest) (*GetCustomActionResponse, error)

	// The Notifications (get) method retrieves the notification periods for the file management
	// job and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+--------------------------------------+
	//	|         RETURN          |                                      |
	//	|       VALUE/CODE        |             DESCRIPTION              |
	//	|                         |                                      |
	//	+-------------------------+--------------------------------------+
	//	+-------------------------+--------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The notifications parameter is NULL. |
	//	+-------------------------+--------------------------------------+
	GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error)

	// Logging operation.
	GetLogging(context.Context, *GetLoggingRequest) (*GetLoggingResponse, error)

	// Logging operation.
	SetLogging(context.Context, *SetLoggingRequest) (*SetLoggingResponse, error)

	// ReportEnabled operation.
	GetReportEnabled(context.Context, *GetReportEnabledRequest) (*GetReportEnabledResponse, error)

	// ReportEnabled operation.
	SetReportEnabled(context.Context, *SetReportEnabledRequest) (*SetReportEnabledResponse, error)

	// Formats operation.
	GetFormats(context.Context, *GetFormatsRequest) (*GetFormatsResponse, error)

	// Formats operation.
	SetFormats(context.Context, *SetFormatsRequest) (*SetFormatsResponse, error)

	// MailTo operation.
	GetMailTo(context.Context, *GetMailToRequest) (*GetMailToResponse, error)

	// MailTo operation.
	SetMailTo(context.Context, *SetMailToRequest) (*SetMailToResponse, error)

	// DaysSinceFileCreated operation.
	GetDaysSinceFileCreated(context.Context, *GetDaysSinceFileCreatedRequest) (*GetDaysSinceFileCreatedResponse, error)

	// DaysSinceFileCreated operation.
	SetDaysSinceFileCreated(context.Context, *SetDaysSinceFileCreatedRequest) (*SetDaysSinceFileCreatedResponse, error)

	// The DaysSinceFileLastAccessed (put) method sets the days since file last accessed
	// property for the file management job.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-----------------------------------------------+
	//	|         RETURN          |                                               |
	//	|       VALUE/CODE        |                  DESCRIPTION                  |
	//	|                         |                                               |
	//	+-------------------------+-----------------------------------------------+
	//	+-------------------------+-----------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The daysSinceAccess parameter is less than 0. |
	//	+-------------------------+-----------------------------------------------+
	GetDaysSinceFileLastAccessed(context.Context, *GetDaysSinceFileLastAccessedRequest) (*GetDaysSinceFileLastAccessedResponse, error)

	// The DaysSinceFileLastAccessed (put) method sets the days since file last accessed
	// property for the file management job.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-----------------------------------------------+
	//	|         RETURN          |                                               |
	//	|       VALUE/CODE        |                  DESCRIPTION                  |
	//	|                         |                                               |
	//	+-------------------------+-----------------------------------------------+
	//	+-------------------------+-----------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The daysSinceAccess parameter is less than 0. |
	//	+-------------------------+-----------------------------------------------+
	SetDaysSinceFileLastAccessed(context.Context, *SetDaysSinceFileLastAccessedRequest) (*SetDaysSinceFileLastAccessedResponse, error)

	// DaysSinceFileLastModified operation.
	GetDaysSinceFileLastModified(context.Context, *GetDaysSinceFileLastModifiedRequest) (*GetDaysSinceFileLastModifiedResponse, error)

	// DaysSinceFileLastModified operation.
	SetDaysSinceFileLastModified(context.Context, *SetDaysSinceFileLastModifiedRequest) (*SetDaysSinceFileLastModifiedResponse, error)

	// The PropertyConditions (get) method retrieves all the property conditions configured
	// for the file management job and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-------------------------------------------+
	//	|         RETURN          |                                           |
	//	|       VALUE/CODE        |                DESCRIPTION                |
	//	|                         |                                           |
	//	+-------------------------+-------------------------------------------+
	//	+-------------------------+-------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The propertyConditions parameter is NULL. |
	//	+-------------------------+-------------------------------------------+
	GetPropertyConditions(context.Context, *GetPropertyConditionsRequest) (*GetPropertyConditionsResponse, error)

	// FromDate operation.
	GetFromDate(context.Context, *GetFromDateRequest) (*GetFromDateResponse, error)

	// FromDate operation.
	SetFromDate(context.Context, *SetFromDateRequest) (*SetFromDateResponse, error)

	// Task operation.
	GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error)

	// Task operation.
	SetTask(context.Context, *SetTaskRequest) (*SetTaskResponse, error)

	// Parameters operation.
	GetParameters(context.Context, *GetParametersRequest) (*GetParametersResponse, error)

	// Parameters operation.
	SetParameters(context.Context, *SetParametersRequest) (*SetParametersResponse, error)

	// RunningStatus operation.
	GetRunningStatus(context.Context, *GetRunningStatusRequest) (*GetRunningStatusResponse, error)

	// LastError operation.
	GetLastError(context.Context, *GetLastErrorRequest) (*GetLastErrorResponse, error)

	// The LastReportPathWithoutExtension (get) method retrieves the last report directory
	// without extension where the generated report(s) was (were) stored when the file management
	// job was previously run and returns S_OK upon successful completion.
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
	GetLastReportPathWithoutExtension(context.Context, *GetLastReportPathWithoutExtensionRequest) (*GetLastReportPathWithoutExtensionResponse, error)

	// LastRun operation.
	GetLastRun(context.Context, *GetLastRunRequest) (*GetLastRunResponse, error)

	// FileNamePattern operation.
	GetFileNamePattern(context.Context, *GetFileNamePatternRequest) (*GetFileNamePatternResponse, error)

	// FileNamePattern operation.
	SetFileNamePattern(context.Context, *SetFileNamePatternRequest) (*SetFileNamePatternResponse, error)

	// Run operation.
	Run(context.Context, *RunRequest) (*RunResponse, error)

	// WaitForCompletion operation.
	WaitForCompletion(context.Context, *WaitForCompletionRequest) (*WaitForCompletionResponse, error)

	// Cancel operation.
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)

	// The AddNotification method adds a notification period to the file management job's
	// list of notification periods.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+----------------------------------------------------------------+
	//	|              RETURN              |                                                                |
	//	|            VALUE/CODE            |                          DESCRIPTION                           |
	//	|                                  |                                                                |
	//	+----------------------------------+----------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------+
	//	| 0x80045303 FSRM_E_ALREADY_EXISTS | The object already exists.                                     |
	//	+----------------------------------+----------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | The days parameter is not a valid value; it is less than zero. |
	//	+----------------------------------+----------------------------------------------------------------+
	AddNotification(context.Context, *AddNotificationRequest) (*AddNotificationResponse, error)

	// The DeleteNotification method deletes a notification period from the file management
	// job's list of notification periods.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+------------------------------------------------+
	//	|           RETURN            |                                                |
	//	|         VALUE/CODE          |                  DESCRIPTION                   |
	//	|                             |                                                |
	//	+-----------------------------+------------------------------------------------+
	//	+-----------------------------+------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | The specified notification could not be found. |
	//	+-----------------------------+------------------------------------------------+
	//
	// If any validation fails, the server MUST terminate processing and return a nonzero
	// error code.
	DeleteNotification(context.Context, *DeleteNotificationRequest) (*DeleteNotificationResponse, error)

	// The ModifyNotification method modifies the value of an existing notification period
	// in the file management job's list of notification periods.
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
	//	| 0x80045301 FSRM_E_ NOT_FOUND     | The specified notification could not be found.                                   |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045303 FSRM_E_ALREADY_EXISTS | The object already exists.                                                       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This code is returned for the following reasons: The days parameter is not a     |
	//	|                                  | valid value; it is less than zero. The newDays parameter is not a valid value;   |
	//	|                                  | it is less than zero.                                                            |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	ModifyNotification(context.Context, *ModifyNotificationRequest) (*ModifyNotificationResponse, error)

	// The CreateNotificationAction method creates a notification and associates it with
	// the specified notification period.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND       | The specified notification could not be found.                                   |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045303 FSRM_E_ ALREADY_EXISTS | The object already exists.                                                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG           | This code is returned for the following reasons: The days parameter is not       |
	//	|                                   | a valid value; it is less than zero. The actionType parameter is not a valid     |
	//	|                                   | value. If actionType is FsrmActionType_Unknown, the parameter MUST be considered |
	//	|                                   | an invalid value. The action parameter is NULL.                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	CreateNotificationAction(context.Context, *CreateNotificationActionRequest) (*CreateNotificationActionResponse, error)

	// The EnumNotificationActions method enumerates all the notifications for the specified
	// notification period.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+------------------------------------------------+
	//	|           RETURN            |                                                |
	//	|         VALUE/CODE          |                  DESCRIPTION                   |
	//	|                             |                                                |
	//	+-----------------------------+------------------------------------------------+
	//	+-----------------------------+------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | The specified notification could not be found. |
	//	+-----------------------------+------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG     | The actions parameter is NULL.                 |
	//	+-----------------------------+------------------------------------------------+
	EnumNotificationActions(context.Context, *EnumNotificationActionsRequest) (*EnumNotificationActionsResponse, error)

	// The CreatePropertyCondition method creates a property condition associated with the
	// file management job.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND       | A property definition with the specified name does not exist.                    |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045303 FSRM_E_ ALREADY_EXISTS | The object already exists.                                                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_ INVALIDARG          | This code is returned for the following reasons: The propertyCondition parameter |
	//	|                                   | is NULL. The name parameter does not conform to the requirements for a property  |
	//	|                                   | definitions name.                                                                |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	CreatePropertyCondition(context.Context, *CreatePropertyConditionRequest) (*CreatePropertyConditionResponse, error)

	// The CreateCustomAction method creates a command line action type notification for
	// the file management job's custom action and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+-------------------------------------+
	//	|              RETURN              |                                     |
	//	|            VALUE/CODE            |             DESCRIPTION             |
	//	|                                  |                                     |
	//	+----------------------------------+-------------------------------------+
	//	+----------------------------------+-------------------------------------+
	//	| 0x80045303 FSMR_E_ALREADY_EXISTS | A custom action already exists.     |
	//	+----------------------------------+-------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | The customAction parameter is NULL. |
	//	+----------------------------------+-------------------------------------+
	CreateCustomAction(context.Context, *CreateCustomActionRequest) (*CreateCustomActionResponse, error)
}

func RegisterFileManagementJobServer(conn dcerpc.Conn, o FileManagementJobServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFileManagementJobServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FileManagementJobSyntaxV1_0))...)
}

func NewFileManagementJobServerHandle(o FileManagementJobServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FileManagementJobServerHandle(ctx, o, opNum, r)
	}
}

func FileManagementJobServerHandle(ctx context.Context, o FileManagementJobServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 12 {
		// IFsrmObject base method.
		return ifsrmobject.ObjectServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 12: // Name
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // Name
		in := &SetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetName(ctx, in)
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
	case 16: // Enabled
		in := &GetEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // Enabled
		in := &SetEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // OperationType
		in := &GetOperationTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOperationType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // OperationType
		in := &SetOperationTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOperationType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // ExpirationDirectory
		in := &GetExpirationDirectoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetExpirationDirectory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // ExpirationDirectory
		in := &SetExpirationDirectoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetExpirationDirectory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // CustomAction
		in := &GetCustomActionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCustomAction(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // Notifications
		in := &GetNotificationsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNotifications(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // Logging
		in := &GetLoggingRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLogging(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // Logging
		in := &SetLoggingRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetLogging(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // ReportEnabled
		in := &GetReportEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetReportEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // ReportEnabled
		in := &SetReportEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetReportEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // Formats
		in := &GetFormatsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFormats(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // Formats
		in := &SetFormatsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFormats(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // MailTo
		in := &GetMailToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMailTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // MailTo
		in := &SetMailToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMailTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // DaysSinceFileCreated
		in := &GetDaysSinceFileCreatedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDaysSinceFileCreated(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // DaysSinceFileCreated
		in := &SetDaysSinceFileCreatedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDaysSinceFileCreated(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // DaysSinceFileLastAccessed
		in := &GetDaysSinceFileLastAccessedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDaysSinceFileLastAccessed(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // DaysSinceFileLastAccessed
		in := &SetDaysSinceFileLastAccessedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDaysSinceFileLastAccessed(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // DaysSinceFileLastModified
		in := &GetDaysSinceFileLastModifiedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDaysSinceFileLastModified(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // DaysSinceFileLastModified
		in := &SetDaysSinceFileLastModifiedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDaysSinceFileLastModified(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // PropertyConditions
		in := &GetPropertyConditionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPropertyConditions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // FromDate
		in := &GetFromDateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFromDate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // FromDate
		in := &SetFromDateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFromDate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // Task
		in := &GetTaskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTask(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // Task
		in := &SetTaskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetTask(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // Parameters
		in := &GetParametersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetParameters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // Parameters
		in := &SetParametersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetParameters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // RunningStatus
		in := &GetRunningStatusRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRunningStatus(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // LastError
		in := &GetLastErrorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLastError(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // LastReportPathWithoutExtension
		in := &GetLastReportPathWithoutExtensionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLastReportPathWithoutExtension(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 48: // LastRun
		in := &GetLastRunRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLastRun(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // FileNamePattern
		in := &GetFileNamePatternRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileNamePattern(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 50: // FileNamePattern
		in := &SetFileNamePatternRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFileNamePattern(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 51: // Run
		in := &RunRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Run(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 52: // WaitForCompletion
		in := &WaitForCompletionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.WaitForCompletion(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 53: // Cancel
		in := &CancelRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Cancel(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 54: // AddNotification
		in := &AddNotificationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddNotification(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 55: // DeleteNotification
		in := &DeleteNotificationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteNotification(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 56: // ModifyNotification
		in := &ModifyNotificationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ModifyNotification(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 57: // CreateNotificationAction
		in := &CreateNotificationActionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateNotificationAction(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 58: // EnumNotificationActions
		in := &EnumNotificationActionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumNotificationActions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 59: // CreatePropertyCondition
		in := &CreatePropertyConditionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreatePropertyCondition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 60: // CreateCustomAction
		in := &CreateCustomActionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateCustomAction(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
