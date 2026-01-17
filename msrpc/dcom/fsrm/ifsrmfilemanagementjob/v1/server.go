package ifsrmfilemanagementjob

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
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Name
		op := &xxx_SetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetName(ctx, req)
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
	case 16: // Enabled
		op := &xxx_GetEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // Enabled
		op := &xxx_SetEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // OperationType
		op := &xxx_GetOperationTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOperationTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOperationType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // OperationType
		op := &xxx_SetOperationTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOperationTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOperationType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // ExpirationDirectory
		op := &xxx_GetExpirationDirectoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetExpirationDirectoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetExpirationDirectory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // ExpirationDirectory
		op := &xxx_SetExpirationDirectoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetExpirationDirectoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetExpirationDirectory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // CustomAction
		op := &xxx_GetCustomActionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCustomActionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCustomAction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // Notifications
		op := &xxx_GetNotificationsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNotificationsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNotifications(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // Logging
		op := &xxx_GetLoggingOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLoggingRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLogging(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // Logging
		op := &xxx_SetLoggingOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLoggingRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLogging(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // ReportEnabled
		op := &xxx_GetReportEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReportEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReportEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // ReportEnabled
		op := &xxx_SetReportEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetReportEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetReportEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // Formats
		op := &xxx_GetFormatsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFormatsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFormats(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // Formats
		op := &xxx_SetFormatsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFormatsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFormats(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // MailTo
		op := &xxx_GetMailToOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMailToRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMailTo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // MailTo
		op := &xxx_SetMailToOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMailToRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMailTo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // DaysSinceFileCreated
		op := &xxx_GetDaysSinceFileCreatedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDaysSinceFileCreatedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDaysSinceFileCreated(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // DaysSinceFileCreated
		op := &xxx_SetDaysSinceFileCreatedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDaysSinceFileCreatedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDaysSinceFileCreated(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // DaysSinceFileLastAccessed
		op := &xxx_GetDaysSinceFileLastAccessedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDaysSinceFileLastAccessedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDaysSinceFileLastAccessed(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // DaysSinceFileLastAccessed
		op := &xxx_SetDaysSinceFileLastAccessedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDaysSinceFileLastAccessedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDaysSinceFileLastAccessed(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // DaysSinceFileLastModified
		op := &xxx_GetDaysSinceFileLastModifiedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDaysSinceFileLastModifiedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDaysSinceFileLastModified(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // DaysSinceFileLastModified
		op := &xxx_SetDaysSinceFileLastModifiedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDaysSinceFileLastModifiedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDaysSinceFileLastModified(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // PropertyConditions
		op := &xxx_GetPropertyConditionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertyConditionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPropertyConditions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // FromDate
		op := &xxx_GetFromDateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFromDateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFromDate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // FromDate
		op := &xxx_SetFromDateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFromDateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFromDate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // Task
		op := &xxx_GetTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // Task
		op := &xxx_SetTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // Parameters
		op := &xxx_GetParametersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetParametersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetParameters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // Parameters
		op := &xxx_SetParametersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetParametersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetParameters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // RunningStatus
		op := &xxx_GetRunningStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRunningStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRunningStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // LastError
		op := &xxx_GetLastErrorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLastErrorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLastError(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // LastReportPathWithoutExtension
		op := &xxx_GetLastReportPathWithoutExtensionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLastReportPathWithoutExtensionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLastReportPathWithoutExtension(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // LastRun
		op := &xxx_GetLastRunOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLastRunRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLastRun(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // FileNamePattern
		op := &xxx_GetFileNamePatternOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileNamePatternRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileNamePattern(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 50: // FileNamePattern
		op := &xxx_SetFileNamePatternOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFileNamePatternRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFileNamePattern(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 51: // Run
		op := &xxx_RunOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RunRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Run(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 52: // WaitForCompletion
		op := &xxx_WaitForCompletionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WaitForCompletionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WaitForCompletion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 53: // Cancel
		op := &xxx_CancelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CancelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Cancel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 54: // AddNotification
		op := &xxx_AddNotificationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddNotificationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddNotification(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 55: // DeleteNotification
		op := &xxx_DeleteNotificationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteNotificationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteNotification(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 56: // ModifyNotification
		op := &xxx_ModifyNotificationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyNotificationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyNotification(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 57: // CreateNotificationAction
		op := &xxx_CreateNotificationActionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateNotificationActionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateNotificationAction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 58: // EnumNotificationActions
		op := &xxx_EnumNotificationActionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumNotificationActionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumNotificationActions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 59: // CreatePropertyCondition
		op := &xxx_CreatePropertyConditionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreatePropertyConditionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreatePropertyCondition(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 60: // CreateCustomAction
		op := &xxx_CreateCustomActionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateCustomActionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateCustomAction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmFileManagementJob
type UnimplementedFileManagementJobServer struct {
	ifsrmobject.UnimplementedObjectServer
}

func (UnimplementedFileManagementJobServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetName(context.Context, *SetNameRequest) (*SetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetNamespaceRoots(context.Context, *GetNamespaceRootsRequest) (*GetNamespaceRootsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetNamespaceRoots(context.Context, *SetNamespaceRootsRequest) (*SetNamespaceRootsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetEnabled(context.Context, *GetEnabledRequest) (*GetEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetEnabled(context.Context, *SetEnabledRequest) (*SetEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetOperationType(context.Context, *GetOperationTypeRequest) (*GetOperationTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetOperationType(context.Context, *SetOperationTypeRequest) (*SetOperationTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetExpirationDirectory(context.Context, *GetExpirationDirectoryRequest) (*GetExpirationDirectoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetExpirationDirectory(context.Context, *SetExpirationDirectoryRequest) (*SetExpirationDirectoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetCustomAction(context.Context, *GetCustomActionRequest) (*GetCustomActionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetLogging(context.Context, *GetLoggingRequest) (*GetLoggingResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetLogging(context.Context, *SetLoggingRequest) (*SetLoggingResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetReportEnabled(context.Context, *GetReportEnabledRequest) (*GetReportEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetReportEnabled(context.Context, *SetReportEnabledRequest) (*SetReportEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetFormats(context.Context, *GetFormatsRequest) (*GetFormatsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetFormats(context.Context, *SetFormatsRequest) (*SetFormatsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetMailTo(context.Context, *GetMailToRequest) (*GetMailToResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetMailTo(context.Context, *SetMailToRequest) (*SetMailToResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetDaysSinceFileCreated(context.Context, *GetDaysSinceFileCreatedRequest) (*GetDaysSinceFileCreatedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetDaysSinceFileCreated(context.Context, *SetDaysSinceFileCreatedRequest) (*SetDaysSinceFileCreatedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetDaysSinceFileLastAccessed(context.Context, *GetDaysSinceFileLastAccessedRequest) (*GetDaysSinceFileLastAccessedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetDaysSinceFileLastAccessed(context.Context, *SetDaysSinceFileLastAccessedRequest) (*SetDaysSinceFileLastAccessedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetDaysSinceFileLastModified(context.Context, *GetDaysSinceFileLastModifiedRequest) (*GetDaysSinceFileLastModifiedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetDaysSinceFileLastModified(context.Context, *SetDaysSinceFileLastModifiedRequest) (*SetDaysSinceFileLastModifiedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetPropertyConditions(context.Context, *GetPropertyConditionsRequest) (*GetPropertyConditionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetFromDate(context.Context, *GetFromDateRequest) (*GetFromDateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetFromDate(context.Context, *SetFromDateRequest) (*SetFromDateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetTask(context.Context, *SetTaskRequest) (*SetTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetParameters(context.Context, *GetParametersRequest) (*GetParametersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetParameters(context.Context, *SetParametersRequest) (*SetParametersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetRunningStatus(context.Context, *GetRunningStatusRequest) (*GetRunningStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetLastError(context.Context, *GetLastErrorRequest) (*GetLastErrorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetLastReportPathWithoutExtension(context.Context, *GetLastReportPathWithoutExtensionRequest) (*GetLastReportPathWithoutExtensionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetLastRun(context.Context, *GetLastRunRequest) (*GetLastRunResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) GetFileNamePattern(context.Context, *GetFileNamePatternRequest) (*GetFileNamePatternResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) SetFileNamePattern(context.Context, *SetFileNamePatternRequest) (*SetFileNamePatternResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) Run(context.Context, *RunRequest) (*RunResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) WaitForCompletion(context.Context, *WaitForCompletionRequest) (*WaitForCompletionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) Cancel(context.Context, *CancelRequest) (*CancelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) AddNotification(context.Context, *AddNotificationRequest) (*AddNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) DeleteNotification(context.Context, *DeleteNotificationRequest) (*DeleteNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) ModifyNotification(context.Context, *ModifyNotificationRequest) (*ModifyNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) CreateNotificationAction(context.Context, *CreateNotificationActionRequest) (*CreateNotificationActionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) EnumNotificationActions(context.Context, *EnumNotificationActionsRequest) (*EnumNotificationActionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) CreatePropertyCondition(context.Context, *CreatePropertyConditionRequest) (*CreatePropertyConditionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileManagementJobServer) CreateCustomAction(context.Context, *CreateCustomActionRequest) (*CreateCustomActionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FileManagementJobServer = (*UnimplementedFileManagementJobServer)(nil)
