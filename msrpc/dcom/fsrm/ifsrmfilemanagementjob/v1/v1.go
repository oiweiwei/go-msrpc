package ifsrmfilemanagementjob

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
	ifsrmobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmobject/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = ifsrmobject.GoPackage
	_ = oaut.GoPackage
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmFileManagementJob interface identifier 0770687e-9f36-4d6f-8778-599d188461c9
	FileManagementJobIID = &dcom.IID{Data1: 0x0770687e, Data2: 0x9f36, Data3: 0x4d6f, Data4: []byte{0x87, 0x78, 0x59, 0x9d, 0x18, 0x84, 0x61, 0xc9}}
	// Syntax UUID
	FileManagementJobSyntaxUUID = &uuid.UUID{TimeLow: 0x770687e, TimeMid: 0x9f36, TimeHiAndVersion: 0x4d6f, ClockSeqHiAndReserved: 0x87, ClockSeqLow: 0x78, Node: [6]uint8{0x59, 0x9d, 0x18, 0x84, 0x61, 0xc9}}
	// Syntax ID
	FileManagementJobSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: FileManagementJobSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// IFsrmFileManagementJob interface.
type FileManagementJobClient interface {

	// IFsrmObject retrieval method.
	Object() ifsrmobject.ObjectClient

	// Name operation.
	GetName(context.Context, *GetNameRequest, ...dcerpc.CallOption) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest, ...dcerpc.CallOption) (*SetNameResponse, error)

	// NamespaceRoots operation.
	GetNamespaceRoots(context.Context, *GetNamespaceRootsRequest, ...dcerpc.CallOption) (*GetNamespaceRootsResponse, error)

	// NamespaceRoots operation.
	SetNamespaceRoots(context.Context, *SetNamespaceRootsRequest, ...dcerpc.CallOption) (*SetNamespaceRootsResponse, error)

	// Enabled operation.
	GetEnabled(context.Context, *GetEnabledRequest, ...dcerpc.CallOption) (*GetEnabledResponse, error)

	// Enabled operation.
	SetEnabled(context.Context, *SetEnabledRequest, ...dcerpc.CallOption) (*SetEnabledResponse, error)

	// OperationType operation.
	GetOperationType(context.Context, *GetOperationTypeRequest, ...dcerpc.CallOption) (*GetOperationTypeResponse, error)

	// OperationType operation.
	SetOperationType(context.Context, *SetOperationTypeRequest, ...dcerpc.CallOption) (*SetOperationTypeResponse, error)

	// ExpirationDirectory operation.
	GetExpirationDirectory(context.Context, *GetExpirationDirectoryRequest, ...dcerpc.CallOption) (*GetExpirationDirectoryResponse, error)

	// ExpirationDirectory operation.
	SetExpirationDirectory(context.Context, *SetExpirationDirectoryRequest, ...dcerpc.CallOption) (*SetExpirationDirectoryResponse, error)

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
	GetCustomAction(context.Context, *GetCustomActionRequest, ...dcerpc.CallOption) (*GetCustomActionResponse, error)

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
	GetNotifications(context.Context, *GetNotificationsRequest, ...dcerpc.CallOption) (*GetNotificationsResponse, error)

	// Logging operation.
	GetLogging(context.Context, *GetLoggingRequest, ...dcerpc.CallOption) (*GetLoggingResponse, error)

	// Logging operation.
	SetLogging(context.Context, *SetLoggingRequest, ...dcerpc.CallOption) (*SetLoggingResponse, error)

	// ReportEnabled operation.
	GetReportEnabled(context.Context, *GetReportEnabledRequest, ...dcerpc.CallOption) (*GetReportEnabledResponse, error)

	// ReportEnabled operation.
	SetReportEnabled(context.Context, *SetReportEnabledRequest, ...dcerpc.CallOption) (*SetReportEnabledResponse, error)

	// Formats operation.
	GetFormats(context.Context, *GetFormatsRequest, ...dcerpc.CallOption) (*GetFormatsResponse, error)

	// Formats operation.
	SetFormats(context.Context, *SetFormatsRequest, ...dcerpc.CallOption) (*SetFormatsResponse, error)

	// MailTo operation.
	GetMailTo(context.Context, *GetMailToRequest, ...dcerpc.CallOption) (*GetMailToResponse, error)

	// MailTo operation.
	SetMailTo(context.Context, *SetMailToRequest, ...dcerpc.CallOption) (*SetMailToResponse, error)

	// DaysSinceFileCreated operation.
	GetDaysSinceFileCreated(context.Context, *GetDaysSinceFileCreatedRequest, ...dcerpc.CallOption) (*GetDaysSinceFileCreatedResponse, error)

	// DaysSinceFileCreated operation.
	SetDaysSinceFileCreated(context.Context, *SetDaysSinceFileCreatedRequest, ...dcerpc.CallOption) (*SetDaysSinceFileCreatedResponse, error)

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
	GetDaysSinceFileLastAccessed(context.Context, *GetDaysSinceFileLastAccessedRequest, ...dcerpc.CallOption) (*GetDaysSinceFileLastAccessedResponse, error)

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
	SetDaysSinceFileLastAccessed(context.Context, *SetDaysSinceFileLastAccessedRequest, ...dcerpc.CallOption) (*SetDaysSinceFileLastAccessedResponse, error)

	// DaysSinceFileLastModified operation.
	GetDaysSinceFileLastModified(context.Context, *GetDaysSinceFileLastModifiedRequest, ...dcerpc.CallOption) (*GetDaysSinceFileLastModifiedResponse, error)

	// DaysSinceFileLastModified operation.
	SetDaysSinceFileLastModified(context.Context, *SetDaysSinceFileLastModifiedRequest, ...dcerpc.CallOption) (*SetDaysSinceFileLastModifiedResponse, error)

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
	GetPropertyConditions(context.Context, *GetPropertyConditionsRequest, ...dcerpc.CallOption) (*GetPropertyConditionsResponse, error)

	// FromDate operation.
	GetFromDate(context.Context, *GetFromDateRequest, ...dcerpc.CallOption) (*GetFromDateResponse, error)

	// FromDate operation.
	SetFromDate(context.Context, *SetFromDateRequest, ...dcerpc.CallOption) (*SetFromDateResponse, error)

	// Task operation.
	GetTask(context.Context, *GetTaskRequest, ...dcerpc.CallOption) (*GetTaskResponse, error)

	// Task operation.
	SetTask(context.Context, *SetTaskRequest, ...dcerpc.CallOption) (*SetTaskResponse, error)

	// Parameters operation.
	GetParameters(context.Context, *GetParametersRequest, ...dcerpc.CallOption) (*GetParametersResponse, error)

	// Parameters operation.
	SetParameters(context.Context, *SetParametersRequest, ...dcerpc.CallOption) (*SetParametersResponse, error)

	// RunningStatus operation.
	GetRunningStatus(context.Context, *GetRunningStatusRequest, ...dcerpc.CallOption) (*GetRunningStatusResponse, error)

	// LastError operation.
	GetLastError(context.Context, *GetLastErrorRequest, ...dcerpc.CallOption) (*GetLastErrorResponse, error)

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
	GetLastReportPathWithoutExtension(context.Context, *GetLastReportPathWithoutExtensionRequest, ...dcerpc.CallOption) (*GetLastReportPathWithoutExtensionResponse, error)

	// LastRun operation.
	GetLastRun(context.Context, *GetLastRunRequest, ...dcerpc.CallOption) (*GetLastRunResponse, error)

	// FileNamePattern operation.
	GetFileNamePattern(context.Context, *GetFileNamePatternRequest, ...dcerpc.CallOption) (*GetFileNamePatternResponse, error)

	// FileNamePattern operation.
	SetFileNamePattern(context.Context, *SetFileNamePatternRequest, ...dcerpc.CallOption) (*SetFileNamePatternResponse, error)

	// Run operation.
	Run(context.Context, *RunRequest, ...dcerpc.CallOption) (*RunResponse, error)

	// WaitForCompletion operation.
	WaitForCompletion(context.Context, *WaitForCompletionRequest, ...dcerpc.CallOption) (*WaitForCompletionResponse, error)

	// Cancel operation.
	Cancel(context.Context, *CancelRequest, ...dcerpc.CallOption) (*CancelResponse, error)

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
	AddNotification(context.Context, *AddNotificationRequest, ...dcerpc.CallOption) (*AddNotificationResponse, error)

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
	DeleteNotification(context.Context, *DeleteNotificationRequest, ...dcerpc.CallOption) (*DeleteNotificationResponse, error)

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
	ModifyNotification(context.Context, *ModifyNotificationRequest, ...dcerpc.CallOption) (*ModifyNotificationResponse, error)

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
	CreateNotificationAction(context.Context, *CreateNotificationActionRequest, ...dcerpc.CallOption) (*CreateNotificationActionResponse, error)

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
	EnumNotificationActions(context.Context, *EnumNotificationActionsRequest, ...dcerpc.CallOption) (*EnumNotificationActionsResponse, error)

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
	CreatePropertyCondition(context.Context, *CreatePropertyConditionRequest, ...dcerpc.CallOption) (*CreatePropertyConditionResponse, error)

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
	CreateCustomAction(context.Context, *CreateCustomActionRequest, ...dcerpc.CallOption) (*CreateCustomActionResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) FileManagementJobClient
}

// DaysNotSpecified represents the FsrmDaysNotSpecified RPC constant
var DaysNotSpecified = -1

// DateNotSpecified represents the FsrmDateNotSpecified RPC constant
var DateNotSpecified = -1

type xxx_DefaultFileManagementJobClient struct {
	ifsrmobject.ObjectClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultFileManagementJobClient) Object() ifsrmobject.ObjectClient {
	return o.ObjectClient
}

func (o *xxx_DefaultFileManagementJobClient) GetName(ctx context.Context, in *GetNameRequest, opts ...dcerpc.CallOption) (*GetNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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

func (o *xxx_DefaultFileManagementJobClient) SetName(ctx context.Context, in *SetNameRequest, opts ...dcerpc.CallOption) (*SetNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetNamespaceRoots(ctx context.Context, in *GetNamespaceRootsRequest, opts ...dcerpc.CallOption) (*GetNamespaceRootsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetNamespaceRootsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) SetNamespaceRoots(ctx context.Context, in *SetNamespaceRootsRequest, opts ...dcerpc.CallOption) (*SetNamespaceRootsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetNamespaceRootsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetEnabled(ctx context.Context, in *GetEnabledRequest, opts ...dcerpc.CallOption) (*GetEnabledResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) SetEnabled(ctx context.Context, in *SetEnabledRequest, opts ...dcerpc.CallOption) (*SetEnabledResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetOperationType(ctx context.Context, in *GetOperationTypeRequest, opts ...dcerpc.CallOption) (*GetOperationTypeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetOperationTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) SetOperationType(ctx context.Context, in *SetOperationTypeRequest, opts ...dcerpc.CallOption) (*SetOperationTypeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetOperationTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetExpirationDirectory(ctx context.Context, in *GetExpirationDirectoryRequest, opts ...dcerpc.CallOption) (*GetExpirationDirectoryResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetExpirationDirectoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) SetExpirationDirectory(ctx context.Context, in *SetExpirationDirectoryRequest, opts ...dcerpc.CallOption) (*SetExpirationDirectoryResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetExpirationDirectoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetCustomAction(ctx context.Context, in *GetCustomActionRequest, opts ...dcerpc.CallOption) (*GetCustomActionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetCustomActionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...dcerpc.CallOption) (*GetNotificationsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetNotificationsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetLogging(ctx context.Context, in *GetLoggingRequest, opts ...dcerpc.CallOption) (*GetLoggingResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetLoggingResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) SetLogging(ctx context.Context, in *SetLoggingRequest, opts ...dcerpc.CallOption) (*SetLoggingResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetLoggingResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetReportEnabled(ctx context.Context, in *GetReportEnabledRequest, opts ...dcerpc.CallOption) (*GetReportEnabledResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetReportEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) SetReportEnabled(ctx context.Context, in *SetReportEnabledRequest, opts ...dcerpc.CallOption) (*SetReportEnabledResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetReportEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetFormats(ctx context.Context, in *GetFormatsRequest, opts ...dcerpc.CallOption) (*GetFormatsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetFormatsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) SetFormats(ctx context.Context, in *SetFormatsRequest, opts ...dcerpc.CallOption) (*SetFormatsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetFormatsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetMailTo(ctx context.Context, in *GetMailToRequest, opts ...dcerpc.CallOption) (*GetMailToResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetMailToResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) SetMailTo(ctx context.Context, in *SetMailToRequest, opts ...dcerpc.CallOption) (*SetMailToResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetMailToResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetDaysSinceFileCreated(ctx context.Context, in *GetDaysSinceFileCreatedRequest, opts ...dcerpc.CallOption) (*GetDaysSinceFileCreatedResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetDaysSinceFileCreatedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) SetDaysSinceFileCreated(ctx context.Context, in *SetDaysSinceFileCreatedRequest, opts ...dcerpc.CallOption) (*SetDaysSinceFileCreatedResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetDaysSinceFileCreatedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetDaysSinceFileLastAccessed(ctx context.Context, in *GetDaysSinceFileLastAccessedRequest, opts ...dcerpc.CallOption) (*GetDaysSinceFileLastAccessedResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetDaysSinceFileLastAccessedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) SetDaysSinceFileLastAccessed(ctx context.Context, in *SetDaysSinceFileLastAccessedRequest, opts ...dcerpc.CallOption) (*SetDaysSinceFileLastAccessedResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetDaysSinceFileLastAccessedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetDaysSinceFileLastModified(ctx context.Context, in *GetDaysSinceFileLastModifiedRequest, opts ...dcerpc.CallOption) (*GetDaysSinceFileLastModifiedResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetDaysSinceFileLastModifiedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) SetDaysSinceFileLastModified(ctx context.Context, in *SetDaysSinceFileLastModifiedRequest, opts ...dcerpc.CallOption) (*SetDaysSinceFileLastModifiedResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetDaysSinceFileLastModifiedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetPropertyConditions(ctx context.Context, in *GetPropertyConditionsRequest, opts ...dcerpc.CallOption) (*GetPropertyConditionsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetPropertyConditionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetFromDate(ctx context.Context, in *GetFromDateRequest, opts ...dcerpc.CallOption) (*GetFromDateResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetFromDateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) SetFromDate(ctx context.Context, in *SetFromDateRequest, opts ...dcerpc.CallOption) (*SetFromDateResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetFromDateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...dcerpc.CallOption) (*GetTaskResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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

func (o *xxx_DefaultFileManagementJobClient) SetTask(ctx context.Context, in *SetTaskRequest, opts ...dcerpc.CallOption) (*SetTaskResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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

func (o *xxx_DefaultFileManagementJobClient) GetParameters(ctx context.Context, in *GetParametersRequest, opts ...dcerpc.CallOption) (*GetParametersResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetParametersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) SetParameters(ctx context.Context, in *SetParametersRequest, opts ...dcerpc.CallOption) (*SetParametersResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetParametersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetRunningStatus(ctx context.Context, in *GetRunningStatusRequest, opts ...dcerpc.CallOption) (*GetRunningStatusResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetRunningStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetLastError(ctx context.Context, in *GetLastErrorRequest, opts ...dcerpc.CallOption) (*GetLastErrorResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetLastErrorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetLastReportPathWithoutExtension(ctx context.Context, in *GetLastReportPathWithoutExtensionRequest, opts ...dcerpc.CallOption) (*GetLastReportPathWithoutExtensionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetLastReportPathWithoutExtensionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetLastRun(ctx context.Context, in *GetLastRunRequest, opts ...dcerpc.CallOption) (*GetLastRunResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetLastRunResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) GetFileNamePattern(ctx context.Context, in *GetFileNamePatternRequest, opts ...dcerpc.CallOption) (*GetFileNamePatternResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetFileNamePatternResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) SetFileNamePattern(ctx context.Context, in *SetFileNamePatternRequest, opts ...dcerpc.CallOption) (*SetFileNamePatternResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetFileNamePatternResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) Run(ctx context.Context, in *RunRequest, opts ...dcerpc.CallOption) (*RunResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &RunResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) WaitForCompletion(ctx context.Context, in *WaitForCompletionRequest, opts ...dcerpc.CallOption) (*WaitForCompletionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &WaitForCompletionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) Cancel(ctx context.Context, in *CancelRequest, opts ...dcerpc.CallOption) (*CancelResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &CancelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) AddNotification(ctx context.Context, in *AddNotificationRequest, opts ...dcerpc.CallOption) (*AddNotificationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &AddNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) DeleteNotification(ctx context.Context, in *DeleteNotificationRequest, opts ...dcerpc.CallOption) (*DeleteNotificationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &DeleteNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) ModifyNotification(ctx context.Context, in *ModifyNotificationRequest, opts ...dcerpc.CallOption) (*ModifyNotificationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &ModifyNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) CreateNotificationAction(ctx context.Context, in *CreateNotificationActionRequest, opts ...dcerpc.CallOption) (*CreateNotificationActionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &CreateNotificationActionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) EnumNotificationActions(ctx context.Context, in *EnumNotificationActionsRequest, opts ...dcerpc.CallOption) (*EnumNotificationActionsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &EnumNotificationActionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) CreatePropertyCondition(ctx context.Context, in *CreatePropertyConditionRequest, opts ...dcerpc.CallOption) (*CreatePropertyConditionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &CreatePropertyConditionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) CreateCustomAction(ctx context.Context, in *CreateCustomActionRequest, opts ...dcerpc.CallOption) (*CreateCustomActionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &CreateCustomActionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultFileManagementJobClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultFileManagementJobClient) IPID(ctx context.Context, ipid *dcom.IPID) FileManagementJobClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultFileManagementJobClient{
		ObjectClient: o.ObjectClient.IPID(ctx, ipid),
		cc:           o.cc,
		ipid:         ipid,
	}
}

func NewFileManagementJobClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (FileManagementJobClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(FileManagementJobSyntaxV1_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmobject.NewObjectClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultFileManagementJobClient{
		ObjectClient: base,
		cc:           cc,
		ipid:         ipid,
	}, nil
}

// xxx_GetNameOperation structure represents the Name operation
type xxx_GetNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNameOperation) OpNum() int { return 12 }

func (o *xxx_GetNameOperation) OpName() string { return "/IFsrmFileManagementJob/v1/Name" }

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

func (o *GetNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNameOperation) *xxx_GetNameOperation {
	if op == nil {
		op = &xxx_GetNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
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

func (o *GetNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNameOperation) *xxx_GetNameOperation {
	if op == nil {
		op = &xxx_GetNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Name = o.Name
	op.Return = o.Return
	return op
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
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNameOperation structure represents the Name operation
type xxx_SetNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNameOperation) OpNum() int { return 13 }

func (o *xxx_SetNameOperation) OpName() string { return "/IFsrmFileManagementJob/v1/Name" }

func (o *xxx_SetNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_SetNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_SetNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNameRequest structure represents the Name operation request
type SetNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Name *oaut.String   `idl:"name:name" json:"name"`
}

func (o *SetNameRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNameOperation) *xxx_SetNameOperation {
	if op == nil {
		op = &xxx_SetNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Name = o.Name
	return op
}

func (o *SetNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
}
func (o *SetNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNameResponse structure represents the Name operation response
type SetNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Name return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNameResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNameOperation) *xxx_SetNameOperation {
	if op == nil {
		op = &xxx_SetNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNamespaceRootsOperation structure represents the NamespaceRoots operation
type xxx_GetNamespaceRootsOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	NamespaceRoots *oaut.SafeArray `idl:"name:namespaceRoots" json:"namespace_roots"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNamespaceRootsOperation) OpNum() int { return 14 }

func (o *xxx_GetNamespaceRootsOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/NamespaceRoots"
}

func (o *xxx_GetNamespaceRootsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNamespaceRootsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetNamespaceRootsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetNamespaceRootsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNamespaceRootsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// namespaceRoots {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.NamespaceRoots != nil {
			_ptr_namespaceRoots := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NamespaceRoots != nil {
					if err := o.NamespaceRoots.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NamespaceRoots, _ptr_namespaceRoots); err != nil {
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

func (o *xxx_GetNamespaceRootsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// namespaceRoots {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_namespaceRoots := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NamespaceRoots == nil {
				o.NamespaceRoots = &oaut.SafeArray{}
			}
			if err := o.NamespaceRoots.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_namespaceRoots := func(ptr interface{}) { o.NamespaceRoots = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.NamespaceRoots, _s_namespaceRoots, _ptr_namespaceRoots); err != nil {
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

// GetNamespaceRootsRequest structure represents the NamespaceRoots operation request
type GetNamespaceRootsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetNamespaceRootsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNamespaceRootsOperation) *xxx_GetNamespaceRootsOperation {
	if op == nil {
		op = &xxx_GetNamespaceRootsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetNamespaceRootsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNamespaceRootsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetNamespaceRootsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNamespaceRootsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNamespaceRootsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNamespaceRootsResponse structure represents the NamespaceRoots operation response
type GetNamespaceRootsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	NamespaceRoots *oaut.SafeArray `idl:"name:namespaceRoots" json:"namespace_roots"`
	// Return: The NamespaceRoots return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNamespaceRootsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNamespaceRootsOperation) *xxx_GetNamespaceRootsOperation {
	if op == nil {
		op = &xxx_GetNamespaceRootsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.NamespaceRoots = o.NamespaceRoots
	op.Return = o.Return
	return op
}

func (o *GetNamespaceRootsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNamespaceRootsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.NamespaceRoots = op.NamespaceRoots
	o.Return = op.Return
}
func (o *GetNamespaceRootsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNamespaceRootsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNamespaceRootsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNamespaceRootsOperation structure represents the NamespaceRoots operation
type xxx_SetNamespaceRootsOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	NamespaceRoots *oaut.SafeArray `idl:"name:namespaceRoots" json:"namespace_roots"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNamespaceRootsOperation) OpNum() int { return 15 }

func (o *xxx_SetNamespaceRootsOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/NamespaceRoots"
}

func (o *xxx_SetNamespaceRootsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNamespaceRootsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// namespaceRoots {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.NamespaceRoots != nil {
			_ptr_namespaceRoots := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NamespaceRoots != nil {
					if err := o.NamespaceRoots.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NamespaceRoots, _ptr_namespaceRoots); err != nil {
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

func (o *xxx_SetNamespaceRootsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// namespaceRoots {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_namespaceRoots := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NamespaceRoots == nil {
				o.NamespaceRoots = &oaut.SafeArray{}
			}
			if err := o.NamespaceRoots.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_namespaceRoots := func(ptr interface{}) { o.NamespaceRoots = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.NamespaceRoots, _s_namespaceRoots, _ptr_namespaceRoots); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNamespaceRootsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNamespaceRootsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNamespaceRootsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNamespaceRootsRequest structure represents the NamespaceRoots operation request
type SetNamespaceRootsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	NamespaceRoots *oaut.SafeArray `idl:"name:namespaceRoots" json:"namespace_roots"`
}

func (o *SetNamespaceRootsRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNamespaceRootsOperation) *xxx_SetNamespaceRootsOperation {
	if op == nil {
		op = &xxx_SetNamespaceRootsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.NamespaceRoots = o.NamespaceRoots
	return op
}

func (o *SetNamespaceRootsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNamespaceRootsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NamespaceRoots = op.NamespaceRoots
}
func (o *SetNamespaceRootsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNamespaceRootsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNamespaceRootsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNamespaceRootsResponse structure represents the NamespaceRoots operation response
type SetNamespaceRootsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The NamespaceRoots return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNamespaceRootsResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNamespaceRootsOperation) *xxx_SetNamespaceRootsOperation {
	if op == nil {
		op = &xxx_SetNamespaceRootsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNamespaceRootsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNamespaceRootsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNamespaceRootsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNamespaceRootsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNamespaceRootsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEnabledOperation structure represents the Enabled operation
type xxx_GetEnabledOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled int16          `idl:"name:enabled" json:"enabled"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEnabledOperation) OpNum() int { return 16 }

func (o *xxx_GetEnabledOperation) OpName() string { return "/IFsrmFileManagementJob/v1/Enabled" }

func (o *xxx_GetEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetEnabledRequest structure represents the Enabled operation request
type GetEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEnabledRequest) xxx_ToOp(ctx context.Context, op *xxx_GetEnabledOperation) *xxx_GetEnabledOperation {
	if op == nil {
		op = &xxx_GetEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEnabledResponse structure represents the Enabled operation response
type GetEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled int16          `idl:"name:enabled" json:"enabled"`
	// Return: The Enabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEnabledResponse) xxx_ToOp(ctx context.Context, op *xxx_GetEnabledOperation) *xxx_GetEnabledOperation {
	if op == nil {
		op = &xxx_GetEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enabled = o.Enabled
	op.Return = o.Return
	return op
}

func (o *GetEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enabled = op.Enabled
	o.Return = op.Return
}
func (o *GetEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetEnabledOperation structure represents the Enabled operation
type xxx_SetEnabledOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled int16          `idl:"name:enabled" json:"enabled"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEnabledOperation) OpNum() int { return 17 }

func (o *xxx_SetEnabledOperation) OpName() string { return "/IFsrmFileManagementJob/v1/Enabled" }

func (o *xxx_SetEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_SetEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetEnabledRequest structure represents the Enabled operation request
type SetEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Enabled int16          `idl:"name:enabled" json:"enabled"`
}

func (o *SetEnabledRequest) xxx_ToOp(ctx context.Context, op *xxx_SetEnabledOperation) *xxx_SetEnabledOperation {
	if op == nil {
		op = &xxx_SetEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Enabled = o.Enabled
	return op
}

func (o *SetEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Enabled = op.Enabled
}
func (o *SetEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEnabledResponse structure represents the Enabled operation response
type SetEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Enabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetEnabledResponse) xxx_ToOp(ctx context.Context, op *xxx_SetEnabledOperation) *xxx_SetEnabledOperation {
	if op == nil {
		op = &xxx_SetEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetOperationTypeOperation structure represents the OperationType operation
type xxx_GetOperationTypeOperation struct {
	This          *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat          `idl:"name:That" json:"that"`
	OperationType fsrm.FileManagementType `idl:"name:operationType" json:"operation_type"`
	Return        int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetOperationTypeOperation) OpNum() int { return 18 }

func (o *xxx_GetOperationTypeOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/OperationType"
}

func (o *xxx_GetOperationTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOperationTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetOperationTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetOperationTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOperationTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// operationType {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmFileManagementType}(enum))
	{
		if err := w.WriteEnum(uint16(o.OperationType)); err != nil {
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

func (o *xxx_GetOperationTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// operationType {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmFileManagementType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.OperationType)); err != nil {
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

// GetOperationTypeRequest structure represents the OperationType operation request
type GetOperationTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetOperationTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetOperationTypeOperation) *xxx_GetOperationTypeOperation {
	if op == nil {
		op = &xxx_GetOperationTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetOperationTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetOperationTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetOperationTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetOperationTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOperationTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetOperationTypeResponse structure represents the OperationType operation response
type GetOperationTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat          `idl:"name:That" json:"that"`
	OperationType fsrm.FileManagementType `idl:"name:operationType" json:"operation_type"`
	// Return: The OperationType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetOperationTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetOperationTypeOperation) *xxx_GetOperationTypeOperation {
	if op == nil {
		op = &xxx_GetOperationTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.OperationType = o.OperationType
	op.Return = o.Return
	return op
}

func (o *GetOperationTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetOperationTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.OperationType = op.OperationType
	o.Return = op.Return
}
func (o *GetOperationTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetOperationTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOperationTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetOperationTypeOperation structure represents the OperationType operation
type xxx_SetOperationTypeOperation struct {
	This          *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat          `idl:"name:That" json:"that"`
	OperationType fsrm.FileManagementType `idl:"name:operationType" json:"operation_type"`
	Return        int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetOperationTypeOperation) OpNum() int { return 19 }

func (o *xxx_SetOperationTypeOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/OperationType"
}

func (o *xxx_SetOperationTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOperationTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// operationType {in} (1:{alias=FsrmFileManagementType}(enum))
	{
		if err := w.WriteEnum(uint16(o.OperationType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOperationTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// operationType {in} (1:{alias=FsrmFileManagementType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.OperationType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOperationTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOperationTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetOperationTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetOperationTypeRequest structure represents the OperationType operation request
type SetOperationTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This          *dcom.ORPCThis          `idl:"name:This" json:"this"`
	OperationType fsrm.FileManagementType `idl:"name:operationType" json:"operation_type"`
}

func (o *SetOperationTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_SetOperationTypeOperation) *xxx_SetOperationTypeOperation {
	if op == nil {
		op = &xxx_SetOperationTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.OperationType = o.OperationType
	return op
}

func (o *SetOperationTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetOperationTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.OperationType = op.OperationType
}
func (o *SetOperationTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetOperationTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOperationTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetOperationTypeResponse structure represents the OperationType operation response
type SetOperationTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The OperationType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetOperationTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_SetOperationTypeOperation) *xxx_SetOperationTypeOperation {
	if op == nil {
		op = &xxx_SetOperationTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetOperationTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetOperationTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetOperationTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetOperationTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOperationTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetExpirationDirectoryOperation structure represents the ExpirationDirectory operation
type xxx_GetExpirationDirectoryOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	ExpirationDirectory *oaut.String   `idl:"name:expirationDirectory" json:"expiration_directory"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetExpirationDirectoryOperation) OpNum() int { return 20 }

func (o *xxx_GetExpirationDirectoryOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/ExpirationDirectory"
}

func (o *xxx_GetExpirationDirectoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExpirationDirectoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetExpirationDirectoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetExpirationDirectoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExpirationDirectoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// expirationDirectory {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ExpirationDirectory != nil {
			_ptr_expirationDirectory := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ExpirationDirectory != nil {
					if err := o.ExpirationDirectory.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExpirationDirectory, _ptr_expirationDirectory); err != nil {
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

func (o *xxx_GetExpirationDirectoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// expirationDirectory {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_expirationDirectory := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ExpirationDirectory == nil {
				o.ExpirationDirectory = &oaut.String{}
			}
			if err := o.ExpirationDirectory.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_expirationDirectory := func(ptr interface{}) { o.ExpirationDirectory = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ExpirationDirectory, _s_expirationDirectory, _ptr_expirationDirectory); err != nil {
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

// GetExpirationDirectoryRequest structure represents the ExpirationDirectory operation request
type GetExpirationDirectoryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetExpirationDirectoryRequest) xxx_ToOp(ctx context.Context, op *xxx_GetExpirationDirectoryOperation) *xxx_GetExpirationDirectoryOperation {
	if op == nil {
		op = &xxx_GetExpirationDirectoryOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetExpirationDirectoryRequest) xxx_FromOp(ctx context.Context, op *xxx_GetExpirationDirectoryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetExpirationDirectoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetExpirationDirectoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExpirationDirectoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetExpirationDirectoryResponse structure represents the ExpirationDirectory operation response
type GetExpirationDirectoryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	ExpirationDirectory *oaut.String   `idl:"name:expirationDirectory" json:"expiration_directory"`
	// Return: The ExpirationDirectory return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetExpirationDirectoryResponse) xxx_ToOp(ctx context.Context, op *xxx_GetExpirationDirectoryOperation) *xxx_GetExpirationDirectoryOperation {
	if op == nil {
		op = &xxx_GetExpirationDirectoryOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ExpirationDirectory = o.ExpirationDirectory
	op.Return = o.Return
	return op
}

func (o *GetExpirationDirectoryResponse) xxx_FromOp(ctx context.Context, op *xxx_GetExpirationDirectoryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ExpirationDirectory = op.ExpirationDirectory
	o.Return = op.Return
}
func (o *GetExpirationDirectoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetExpirationDirectoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExpirationDirectoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetExpirationDirectoryOperation structure represents the ExpirationDirectory operation
type xxx_SetExpirationDirectoryOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	ExpirationDirectory *oaut.String   `idl:"name:expirationDirectory" json:"expiration_directory"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetExpirationDirectoryOperation) OpNum() int { return 21 }

func (o *xxx_SetExpirationDirectoryOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/ExpirationDirectory"
}

func (o *xxx_SetExpirationDirectoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExpirationDirectoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// expirationDirectory {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ExpirationDirectory != nil {
			_ptr_expirationDirectory := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ExpirationDirectory != nil {
					if err := o.ExpirationDirectory.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExpirationDirectory, _ptr_expirationDirectory); err != nil {
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

func (o *xxx_SetExpirationDirectoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// expirationDirectory {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_expirationDirectory := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ExpirationDirectory == nil {
				o.ExpirationDirectory = &oaut.String{}
			}
			if err := o.ExpirationDirectory.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_expirationDirectory := func(ptr interface{}) { o.ExpirationDirectory = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ExpirationDirectory, _s_expirationDirectory, _ptr_expirationDirectory); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExpirationDirectoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExpirationDirectoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetExpirationDirectoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetExpirationDirectoryRequest structure represents the ExpirationDirectory operation request
type SetExpirationDirectoryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	ExpirationDirectory *oaut.String   `idl:"name:expirationDirectory" json:"expiration_directory"`
}

func (o *SetExpirationDirectoryRequest) xxx_ToOp(ctx context.Context, op *xxx_SetExpirationDirectoryOperation) *xxx_SetExpirationDirectoryOperation {
	if op == nil {
		op = &xxx_SetExpirationDirectoryOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ExpirationDirectory = o.ExpirationDirectory
	return op
}

func (o *SetExpirationDirectoryRequest) xxx_FromOp(ctx context.Context, op *xxx_SetExpirationDirectoryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ExpirationDirectory = op.ExpirationDirectory
}
func (o *SetExpirationDirectoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetExpirationDirectoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExpirationDirectoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetExpirationDirectoryResponse structure represents the ExpirationDirectory operation response
type SetExpirationDirectoryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ExpirationDirectory return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetExpirationDirectoryResponse) xxx_ToOp(ctx context.Context, op *xxx_SetExpirationDirectoryOperation) *xxx_SetExpirationDirectoryOperation {
	if op == nil {
		op = &xxx_SetExpirationDirectoryOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetExpirationDirectoryResponse) xxx_FromOp(ctx context.Context, op *xxx_SetExpirationDirectoryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetExpirationDirectoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetExpirationDirectoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExpirationDirectoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCustomActionOperation structure represents the CustomAction operation
type xxx_GetCustomActionOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Action *fsrm.ActionCommand `idl:"name:action" json:"action"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCustomActionOperation) OpNum() int { return 22 }

func (o *xxx_GetCustomActionOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/CustomAction"
}

func (o *xxx_GetCustomActionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCustomActionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCustomActionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCustomActionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCustomActionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// action {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmActionCommand}(interface))
	{
		if o.Action != nil {
			_ptr_action := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Action != nil {
					if err := o.Action.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.ActionCommand{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Action, _ptr_action); err != nil {
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

func (o *xxx_GetCustomActionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// action {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmActionCommand}(interface))
	{
		_ptr_action := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Action == nil {
				o.Action = &fsrm.ActionCommand{}
			}
			if err := o.Action.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_action := func(ptr interface{}) { o.Action = *ptr.(**fsrm.ActionCommand) }
		if err := w.ReadPointer(&o.Action, _s_action, _ptr_action); err != nil {
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

// GetCustomActionRequest structure represents the CustomAction operation request
type GetCustomActionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCustomActionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCustomActionOperation) *xxx_GetCustomActionOperation {
	if op == nil {
		op = &xxx_GetCustomActionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetCustomActionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCustomActionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCustomActionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCustomActionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCustomActionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCustomActionResponse structure represents the CustomAction operation response
type GetCustomActionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// action: Pointer to an IFsrmActionCommand interface pointer (section 3.2.4.2.9) that
	// upon completion contains the custom action of the file management job.
	Action *fsrm.ActionCommand `idl:"name:action" json:"action"`
	// Return: The CustomAction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCustomActionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCustomActionOperation) *xxx_GetCustomActionOperation {
	if op == nil {
		op = &xxx_GetCustomActionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Action = o.Action
	op.Return = o.Return
	return op
}

func (o *GetCustomActionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCustomActionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Action = op.Action
	o.Return = op.Return
}
func (o *GetCustomActionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCustomActionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCustomActionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNotificationsOperation structure represents the Notifications operation
type xxx_GetNotificationsOperation struct {
	This          *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Notifications *oaut.SafeArray `idl:"name:notifications" json:"notifications"`
	Return        int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNotificationsOperation) OpNum() int { return 23 }

func (o *xxx_GetNotificationsOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/Notifications"
}

func (o *xxx_GetNotificationsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotificationsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetNotificationsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetNotificationsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotificationsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// notifications {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Notifications != nil {
			_ptr_notifications := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Notifications != nil {
					if err := o.Notifications.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Notifications, _ptr_notifications); err != nil {
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

func (o *xxx_GetNotificationsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// notifications {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_notifications := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Notifications == nil {
				o.Notifications = &oaut.SafeArray{}
			}
			if err := o.Notifications.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_notifications := func(ptr interface{}) { o.Notifications = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Notifications, _s_notifications, _ptr_notifications); err != nil {
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

// GetNotificationsRequest structure represents the Notifications operation request
type GetNotificationsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetNotificationsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNotificationsOperation) *xxx_GetNotificationsOperation {
	if op == nil {
		op = &xxx_GetNotificationsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetNotificationsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNotificationsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetNotificationsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNotificationsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNotificationsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNotificationsResponse structure represents the Notifications operation response
type GetNotificationsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// notifications: Pointer to a SAFEARRAY that upon completion contains all the notification
	// period values for the file management job. A caller MUST release the SAFEARRAY received
	// when the caller is done with it.
	Notifications *oaut.SafeArray `idl:"name:notifications" json:"notifications"`
	// Return: The Notifications return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNotificationsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNotificationsOperation) *xxx_GetNotificationsOperation {
	if op == nil {
		op = &xxx_GetNotificationsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Notifications = o.Notifications
	op.Return = o.Return
	return op
}

func (o *GetNotificationsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNotificationsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Notifications = op.Notifications
	o.Return = op.Return
}
func (o *GetNotificationsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNotificationsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNotificationsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLoggingOperation structure represents the Logging operation
type xxx_GetLoggingOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	LoggingFlags int32          `idl:"name:loggingFlags" json:"logging_flags"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLoggingOperation) OpNum() int { return 24 }

func (o *xxx_GetLoggingOperation) OpName() string { return "/IFsrmFileManagementJob/v1/Logging" }

func (o *xxx_GetLoggingOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLoggingOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLoggingOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLoggingOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLoggingOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// loggingFlags {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.LoggingFlags); err != nil {
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

func (o *xxx_GetLoggingOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// loggingFlags {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.LoggingFlags); err != nil {
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

// GetLoggingRequest structure represents the Logging operation request
type GetLoggingRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLoggingRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLoggingOperation) *xxx_GetLoggingOperation {
	if op == nil {
		op = &xxx_GetLoggingOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetLoggingRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLoggingOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLoggingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLoggingRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLoggingOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLoggingResponse structure represents the Logging operation response
type GetLoggingResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	LoggingFlags int32          `idl:"name:loggingFlags" json:"logging_flags"`
	// Return: The Logging return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLoggingResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLoggingOperation) *xxx_GetLoggingOperation {
	if op == nil {
		op = &xxx_GetLoggingOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.LoggingFlags = o.LoggingFlags
	op.Return = o.Return
	return op
}

func (o *GetLoggingResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLoggingOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LoggingFlags = op.LoggingFlags
	o.Return = op.Return
}
func (o *GetLoggingResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLoggingResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLoggingOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetLoggingOperation structure represents the Logging operation
type xxx_SetLoggingOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	LoggingFlags int32          `idl:"name:loggingFlags" json:"logging_flags"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetLoggingOperation) OpNum() int { return 25 }

func (o *xxx_SetLoggingOperation) OpName() string { return "/IFsrmFileManagementJob/v1/Logging" }

func (o *xxx_SetLoggingOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLoggingOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// loggingFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.LoggingFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLoggingOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// loggingFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.LoggingFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLoggingOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLoggingOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetLoggingOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetLoggingRequest structure represents the Logging operation request
type SetLoggingRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	LoggingFlags int32          `idl:"name:loggingFlags" json:"logging_flags"`
}

func (o *SetLoggingRequest) xxx_ToOp(ctx context.Context, op *xxx_SetLoggingOperation) *xxx_SetLoggingOperation {
	if op == nil {
		op = &xxx_SetLoggingOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LoggingFlags = o.LoggingFlags
	return op
}

func (o *SetLoggingRequest) xxx_FromOp(ctx context.Context, op *xxx_SetLoggingOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LoggingFlags = op.LoggingFlags
}
func (o *SetLoggingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetLoggingRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLoggingOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetLoggingResponse structure represents the Logging operation response
type SetLoggingResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Logging return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetLoggingResponse) xxx_ToOp(ctx context.Context, op *xxx_SetLoggingOperation) *xxx_SetLoggingOperation {
	if op == nil {
		op = &xxx_SetLoggingOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetLoggingResponse) xxx_FromOp(ctx context.Context, op *xxx_SetLoggingOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetLoggingResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetLoggingResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLoggingOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetReportEnabledOperation structure represents the ReportEnabled operation
type xxx_GetReportEnabledOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReportEnabled int16          `idl:"name:reportEnabled" json:"report_enabled"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetReportEnabledOperation) OpNum() int { return 26 }

func (o *xxx_GetReportEnabledOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/ReportEnabled"
}

func (o *xxx_GetReportEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetReportEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetReportEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// reportEnabled {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.ReportEnabled); err != nil {
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

func (o *xxx_GetReportEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// reportEnabled {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.ReportEnabled); err != nil {
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

// GetReportEnabledRequest structure represents the ReportEnabled operation request
type GetReportEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetReportEnabledRequest) xxx_ToOp(ctx context.Context, op *xxx_GetReportEnabledOperation) *xxx_GetReportEnabledOperation {
	if op == nil {
		op = &xxx_GetReportEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetReportEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_GetReportEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetReportEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetReportEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReportEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetReportEnabledResponse structure represents the ReportEnabled operation response
type GetReportEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReportEnabled int16          `idl:"name:reportEnabled" json:"report_enabled"`
	// Return: The ReportEnabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetReportEnabledResponse) xxx_ToOp(ctx context.Context, op *xxx_GetReportEnabledOperation) *xxx_GetReportEnabledOperation {
	if op == nil {
		op = &xxx_GetReportEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReportEnabled = o.ReportEnabled
	op.Return = o.Return
	return op
}

func (o *GetReportEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_GetReportEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReportEnabled = op.ReportEnabled
	o.Return = op.Return
}
func (o *GetReportEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetReportEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReportEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetReportEnabledOperation structure represents the ReportEnabled operation
type xxx_SetReportEnabledOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReportEnabled int16          `idl:"name:reportEnabled" json:"report_enabled"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetReportEnabledOperation) OpNum() int { return 27 }

func (o *xxx_SetReportEnabledOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/ReportEnabled"
}

func (o *xxx_SetReportEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// reportEnabled {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.ReportEnabled); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// reportEnabled {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.ReportEnabled); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetReportEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetReportEnabledRequest structure represents the ReportEnabled operation request
type SetReportEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	ReportEnabled int16          `idl:"name:reportEnabled" json:"report_enabled"`
}

func (o *SetReportEnabledRequest) xxx_ToOp(ctx context.Context, op *xxx_SetReportEnabledOperation) *xxx_SetReportEnabledOperation {
	if op == nil {
		op = &xxx_SetReportEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ReportEnabled = o.ReportEnabled
	return op
}

func (o *SetReportEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_SetReportEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ReportEnabled = op.ReportEnabled
}
func (o *SetReportEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetReportEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetReportEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetReportEnabledResponse structure represents the ReportEnabled operation response
type SetReportEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ReportEnabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetReportEnabledResponse) xxx_ToOp(ctx context.Context, op *xxx_SetReportEnabledOperation) *xxx_SetReportEnabledOperation {
	if op == nil {
		op = &xxx_SetReportEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetReportEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_SetReportEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetReportEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetReportEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetReportEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFormatsOperation structure represents the Formats operation
type xxx_GetFormatsOperation struct {
	This    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Formats *oaut.SafeArray `idl:"name:formats" json:"formats"`
	Return  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFormatsOperation) OpNum() int { return 28 }

func (o *xxx_GetFormatsOperation) OpName() string { return "/IFsrmFileManagementJob/v1/Formats" }

func (o *xxx_GetFormatsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFormatsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFormatsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFormatsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFormatsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// formats {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Formats != nil {
			_ptr_formats := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Formats != nil {
					if err := o.Formats.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Formats, _ptr_formats); err != nil {
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

func (o *xxx_GetFormatsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// formats {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_formats := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Formats == nil {
				o.Formats = &oaut.SafeArray{}
			}
			if err := o.Formats.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_formats := func(ptr interface{}) { o.Formats = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Formats, _s_formats, _ptr_formats); err != nil {
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

// GetFormatsRequest structure represents the Formats operation request
type GetFormatsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFormatsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFormatsOperation) *xxx_GetFormatsOperation {
	if op == nil {
		op = &xxx_GetFormatsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetFormatsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFormatsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFormatsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFormatsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFormatsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFormatsResponse structure represents the Formats operation response
type GetFormatsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Formats *oaut.SafeArray `idl:"name:formats" json:"formats"`
	// Return: The Formats return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFormatsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFormatsOperation) *xxx_GetFormatsOperation {
	if op == nil {
		op = &xxx_GetFormatsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Formats = o.Formats
	op.Return = o.Return
	return op
}

func (o *GetFormatsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFormatsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Formats = op.Formats
	o.Return = op.Return
}
func (o *GetFormatsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFormatsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFormatsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFormatsOperation structure represents the Formats operation
type xxx_SetFormatsOperation struct {
	This    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Formats *oaut.SafeArray `idl:"name:formats" json:"formats"`
	Return  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFormatsOperation) OpNum() int { return 29 }

func (o *xxx_SetFormatsOperation) OpName() string { return "/IFsrmFileManagementJob/v1/Formats" }

func (o *xxx_SetFormatsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFormatsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// formats {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Formats != nil {
			_ptr_formats := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Formats != nil {
					if err := o.Formats.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Formats, _ptr_formats); err != nil {
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

func (o *xxx_SetFormatsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// formats {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_formats := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Formats == nil {
				o.Formats = &oaut.SafeArray{}
			}
			if err := o.Formats.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_formats := func(ptr interface{}) { o.Formats = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Formats, _s_formats, _ptr_formats); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFormatsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFormatsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFormatsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFormatsRequest structure represents the Formats operation request
type SetFormatsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	Formats *oaut.SafeArray `idl:"name:formats" json:"formats"`
}

func (o *SetFormatsRequest) xxx_ToOp(ctx context.Context, op *xxx_SetFormatsOperation) *xxx_SetFormatsOperation {
	if op == nil {
		op = &xxx_SetFormatsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Formats = o.Formats
	return op
}

func (o *SetFormatsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFormatsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Formats = op.Formats
}
func (o *SetFormatsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetFormatsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFormatsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFormatsResponse structure represents the Formats operation response
type SetFormatsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Formats return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFormatsResponse) xxx_ToOp(ctx context.Context, op *xxx_SetFormatsOperation) *xxx_SetFormatsOperation {
	if op == nil {
		op = &xxx_SetFormatsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetFormatsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFormatsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFormatsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetFormatsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFormatsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMailToOperation structure represents the MailTo operation
type xxx_GetMailToOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailTo *oaut.String   `idl:"name:mailTo" json:"mail_to"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMailToOperation) OpNum() int { return 30 }

func (o *xxx_GetMailToOperation) OpName() string { return "/IFsrmFileManagementJob/v1/MailTo" }

func (o *xxx_GetMailToOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailToOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMailToOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMailToOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailToOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// mailTo {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailTo != nil {
			_ptr_mailTo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailTo != nil {
					if err := o.MailTo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailTo, _ptr_mailTo); err != nil {
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

func (o *xxx_GetMailToOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// mailTo {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailTo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailTo == nil {
				o.MailTo = &oaut.String{}
			}
			if err := o.MailTo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailTo := func(ptr interface{}) { o.MailTo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailTo, _s_mailTo, _ptr_mailTo); err != nil {
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

// GetMailToRequest structure represents the MailTo operation request
type GetMailToRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMailToRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMailToOperation) *xxx_GetMailToOperation {
	if op == nil {
		op = &xxx_GetMailToOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMailToRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMailToOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMailToRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMailToRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailToOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMailToResponse structure represents the MailTo operation response
type GetMailToResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailTo *oaut.String   `idl:"name:mailTo" json:"mail_to"`
	// Return: The MailTo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMailToResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMailToOperation) *xxx_GetMailToOperation {
	if op == nil {
		op = &xxx_GetMailToOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MailTo = o.MailTo
	op.Return = o.Return
	return op
}

func (o *GetMailToResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMailToOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MailTo = op.MailTo
	o.Return = op.Return
}
func (o *GetMailToResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMailToResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailToOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMailToOperation structure represents the MailTo operation
type xxx_SetMailToOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailTo *oaut.String   `idl:"name:mailTo" json:"mail_to"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMailToOperation) OpNum() int { return 31 }

func (o *xxx_SetMailToOperation) OpName() string { return "/IFsrmFileManagementJob/v1/MailTo" }

func (o *xxx_SetMailToOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailToOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// mailTo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailTo != nil {
			_ptr_mailTo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailTo != nil {
					if err := o.MailTo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailTo, _ptr_mailTo); err != nil {
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

func (o *xxx_SetMailToOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// mailTo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailTo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailTo == nil {
				o.MailTo = &oaut.String{}
			}
			if err := o.MailTo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailTo := func(ptr interface{}) { o.MailTo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailTo, _s_mailTo, _ptr_mailTo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailToOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailToOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMailToOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMailToRequest structure represents the MailTo operation request
type SetMailToRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	MailTo *oaut.String   `idl:"name:mailTo" json:"mail_to"`
}

func (o *SetMailToRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMailToOperation) *xxx_SetMailToOperation {
	if op == nil {
		op = &xxx_SetMailToOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MailTo = o.MailTo
	return op
}

func (o *SetMailToRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMailToOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MailTo = op.MailTo
}
func (o *SetMailToRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMailToRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailToOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMailToResponse structure represents the MailTo operation response
type SetMailToResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MailTo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMailToResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMailToOperation) *xxx_SetMailToOperation {
	if op == nil {
		op = &xxx_SetMailToOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMailToResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMailToOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMailToResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMailToResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailToOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDaysSinceFileCreatedOperation structure represents the DaysSinceFileCreated operation
type xxx_GetDaysSinceFileCreatedOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	DaysSinceCreation int32          `idl:"name:daysSinceCreation" json:"days_since_creation"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDaysSinceFileCreatedOperation) OpNum() int { return 32 }

func (o *xxx_GetDaysSinceFileCreatedOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/DaysSinceFileCreated"
}

func (o *xxx_GetDaysSinceFileCreatedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDaysSinceFileCreatedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDaysSinceFileCreatedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDaysSinceFileCreatedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDaysSinceFileCreatedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// daysSinceCreation {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.DaysSinceCreation); err != nil {
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

func (o *xxx_GetDaysSinceFileCreatedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// daysSinceCreation {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.DaysSinceCreation); err != nil {
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

// GetDaysSinceFileCreatedRequest structure represents the DaysSinceFileCreated operation request
type GetDaysSinceFileCreatedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDaysSinceFileCreatedRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDaysSinceFileCreatedOperation) *xxx_GetDaysSinceFileCreatedOperation {
	if op == nil {
		op = &xxx_GetDaysSinceFileCreatedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetDaysSinceFileCreatedRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDaysSinceFileCreatedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDaysSinceFileCreatedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDaysSinceFileCreatedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDaysSinceFileCreatedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDaysSinceFileCreatedResponse structure represents the DaysSinceFileCreated operation response
type GetDaysSinceFileCreatedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	DaysSinceCreation int32          `idl:"name:daysSinceCreation" json:"days_since_creation"`
	// Return: The DaysSinceFileCreated return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDaysSinceFileCreatedResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDaysSinceFileCreatedOperation) *xxx_GetDaysSinceFileCreatedOperation {
	if op == nil {
		op = &xxx_GetDaysSinceFileCreatedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DaysSinceCreation = o.DaysSinceCreation
	op.Return = o.Return
	return op
}

func (o *GetDaysSinceFileCreatedResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDaysSinceFileCreatedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DaysSinceCreation = op.DaysSinceCreation
	o.Return = op.Return
}
func (o *GetDaysSinceFileCreatedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDaysSinceFileCreatedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDaysSinceFileCreatedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDaysSinceFileCreatedOperation structure represents the DaysSinceFileCreated operation
type xxx_SetDaysSinceFileCreatedOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	DaysSinceCreation int32          `idl:"name:daysSinceCreation" json:"days_since_creation"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDaysSinceFileCreatedOperation) OpNum() int { return 33 }

func (o *xxx_SetDaysSinceFileCreatedOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/DaysSinceFileCreated"
}

func (o *xxx_SetDaysSinceFileCreatedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDaysSinceFileCreatedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// daysSinceCreation {in} (1:(int32))
	{
		if err := w.WriteData(o.DaysSinceCreation); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDaysSinceFileCreatedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// daysSinceCreation {in} (1:(int32))
	{
		if err := w.ReadData(&o.DaysSinceCreation); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDaysSinceFileCreatedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDaysSinceFileCreatedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetDaysSinceFileCreatedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetDaysSinceFileCreatedRequest structure represents the DaysSinceFileCreated operation request
type SetDaysSinceFileCreatedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	DaysSinceCreation int32          `idl:"name:daysSinceCreation" json:"days_since_creation"`
}

func (o *SetDaysSinceFileCreatedRequest) xxx_ToOp(ctx context.Context, op *xxx_SetDaysSinceFileCreatedOperation) *xxx_SetDaysSinceFileCreatedOperation {
	if op == nil {
		op = &xxx_SetDaysSinceFileCreatedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DaysSinceCreation = o.DaysSinceCreation
	return op
}

func (o *SetDaysSinceFileCreatedRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDaysSinceFileCreatedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DaysSinceCreation = op.DaysSinceCreation
}
func (o *SetDaysSinceFileCreatedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetDaysSinceFileCreatedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDaysSinceFileCreatedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDaysSinceFileCreatedResponse structure represents the DaysSinceFileCreated operation response
type SetDaysSinceFileCreatedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DaysSinceFileCreated return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDaysSinceFileCreatedResponse) xxx_ToOp(ctx context.Context, op *xxx_SetDaysSinceFileCreatedOperation) *xxx_SetDaysSinceFileCreatedOperation {
	if op == nil {
		op = &xxx_SetDaysSinceFileCreatedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetDaysSinceFileCreatedResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDaysSinceFileCreatedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDaysSinceFileCreatedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetDaysSinceFileCreatedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDaysSinceFileCreatedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDaysSinceFileLastAccessedOperation structure represents the DaysSinceFileLastAccessed operation
type xxx_GetDaysSinceFileLastAccessedOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	DaysSinceAccess int32          `idl:"name:daysSinceAccess" json:"days_since_access"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDaysSinceFileLastAccessedOperation) OpNum() int { return 34 }

func (o *xxx_GetDaysSinceFileLastAccessedOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/DaysSinceFileLastAccessed"
}

func (o *xxx_GetDaysSinceFileLastAccessedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDaysSinceFileLastAccessedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDaysSinceFileLastAccessedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDaysSinceFileLastAccessedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDaysSinceFileLastAccessedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// daysSinceAccess {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.DaysSinceAccess); err != nil {
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

func (o *xxx_GetDaysSinceFileLastAccessedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// daysSinceAccess {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.DaysSinceAccess); err != nil {
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

// GetDaysSinceFileLastAccessedRequest structure represents the DaysSinceFileLastAccessed operation request
type GetDaysSinceFileLastAccessedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDaysSinceFileLastAccessedRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDaysSinceFileLastAccessedOperation) *xxx_GetDaysSinceFileLastAccessedOperation {
	if op == nil {
		op = &xxx_GetDaysSinceFileLastAccessedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetDaysSinceFileLastAccessedRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDaysSinceFileLastAccessedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDaysSinceFileLastAccessedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDaysSinceFileLastAccessedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDaysSinceFileLastAccessedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDaysSinceFileLastAccessedResponse structure represents the DaysSinceFileLastAccessed operation response
type GetDaysSinceFileLastAccessedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// daysSinceAccess: Contains the days since file last accessed property for this file
	// management job.
	DaysSinceAccess int32 `idl:"name:daysSinceAccess" json:"days_since_access"`
	// Return: The DaysSinceFileLastAccessed return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDaysSinceFileLastAccessedResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDaysSinceFileLastAccessedOperation) *xxx_GetDaysSinceFileLastAccessedOperation {
	if op == nil {
		op = &xxx_GetDaysSinceFileLastAccessedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DaysSinceAccess = o.DaysSinceAccess
	op.Return = o.Return
	return op
}

func (o *GetDaysSinceFileLastAccessedResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDaysSinceFileLastAccessedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DaysSinceAccess = op.DaysSinceAccess
	o.Return = op.Return
}
func (o *GetDaysSinceFileLastAccessedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDaysSinceFileLastAccessedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDaysSinceFileLastAccessedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDaysSinceFileLastAccessedOperation structure represents the DaysSinceFileLastAccessed operation
type xxx_SetDaysSinceFileLastAccessedOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	DaysSinceAccess int32          `idl:"name:daysSinceAccess" json:"days_since_access"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDaysSinceFileLastAccessedOperation) OpNum() int { return 35 }

func (o *xxx_SetDaysSinceFileLastAccessedOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/DaysSinceFileLastAccessed"
}

func (o *xxx_SetDaysSinceFileLastAccessedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDaysSinceFileLastAccessedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// daysSinceAccess {in} (1:(int32))
	{
		if err := w.WriteData(o.DaysSinceAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDaysSinceFileLastAccessedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// daysSinceAccess {in} (1:(int32))
	{
		if err := w.ReadData(&o.DaysSinceAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDaysSinceFileLastAccessedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDaysSinceFileLastAccessedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetDaysSinceFileLastAccessedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetDaysSinceFileLastAccessedRequest structure represents the DaysSinceFileLastAccessed operation request
type SetDaysSinceFileLastAccessedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// daysSinceAccess: Contains the days since file last accessed property for this file
	// management job.
	DaysSinceAccess int32 `idl:"name:daysSinceAccess" json:"days_since_access"`
}

func (o *SetDaysSinceFileLastAccessedRequest) xxx_ToOp(ctx context.Context, op *xxx_SetDaysSinceFileLastAccessedOperation) *xxx_SetDaysSinceFileLastAccessedOperation {
	if op == nil {
		op = &xxx_SetDaysSinceFileLastAccessedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DaysSinceAccess = o.DaysSinceAccess
	return op
}

func (o *SetDaysSinceFileLastAccessedRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDaysSinceFileLastAccessedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DaysSinceAccess = op.DaysSinceAccess
}
func (o *SetDaysSinceFileLastAccessedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetDaysSinceFileLastAccessedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDaysSinceFileLastAccessedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDaysSinceFileLastAccessedResponse structure represents the DaysSinceFileLastAccessed operation response
type SetDaysSinceFileLastAccessedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DaysSinceFileLastAccessed return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDaysSinceFileLastAccessedResponse) xxx_ToOp(ctx context.Context, op *xxx_SetDaysSinceFileLastAccessedOperation) *xxx_SetDaysSinceFileLastAccessedOperation {
	if op == nil {
		op = &xxx_SetDaysSinceFileLastAccessedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetDaysSinceFileLastAccessedResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDaysSinceFileLastAccessedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDaysSinceFileLastAccessedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetDaysSinceFileLastAccessedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDaysSinceFileLastAccessedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDaysSinceFileLastModifiedOperation structure represents the DaysSinceFileLastModified operation
type xxx_GetDaysSinceFileLastModifiedOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	DaysSinceModify int32          `idl:"name:daysSinceModify" json:"days_since_modify"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDaysSinceFileLastModifiedOperation) OpNum() int { return 36 }

func (o *xxx_GetDaysSinceFileLastModifiedOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/DaysSinceFileLastModified"
}

func (o *xxx_GetDaysSinceFileLastModifiedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDaysSinceFileLastModifiedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDaysSinceFileLastModifiedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDaysSinceFileLastModifiedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDaysSinceFileLastModifiedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// daysSinceModify {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.DaysSinceModify); err != nil {
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

func (o *xxx_GetDaysSinceFileLastModifiedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// daysSinceModify {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.DaysSinceModify); err != nil {
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

// GetDaysSinceFileLastModifiedRequest structure represents the DaysSinceFileLastModified operation request
type GetDaysSinceFileLastModifiedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDaysSinceFileLastModifiedRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDaysSinceFileLastModifiedOperation) *xxx_GetDaysSinceFileLastModifiedOperation {
	if op == nil {
		op = &xxx_GetDaysSinceFileLastModifiedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetDaysSinceFileLastModifiedRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDaysSinceFileLastModifiedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDaysSinceFileLastModifiedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDaysSinceFileLastModifiedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDaysSinceFileLastModifiedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDaysSinceFileLastModifiedResponse structure represents the DaysSinceFileLastModified operation response
type GetDaysSinceFileLastModifiedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	DaysSinceModify int32          `idl:"name:daysSinceModify" json:"days_since_modify"`
	// Return: The DaysSinceFileLastModified return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDaysSinceFileLastModifiedResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDaysSinceFileLastModifiedOperation) *xxx_GetDaysSinceFileLastModifiedOperation {
	if op == nil {
		op = &xxx_GetDaysSinceFileLastModifiedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DaysSinceModify = o.DaysSinceModify
	op.Return = o.Return
	return op
}

func (o *GetDaysSinceFileLastModifiedResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDaysSinceFileLastModifiedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DaysSinceModify = op.DaysSinceModify
	o.Return = op.Return
}
func (o *GetDaysSinceFileLastModifiedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDaysSinceFileLastModifiedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDaysSinceFileLastModifiedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDaysSinceFileLastModifiedOperation structure represents the DaysSinceFileLastModified operation
type xxx_SetDaysSinceFileLastModifiedOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	DaysSinceModify int32          `idl:"name:daysSinceModify" json:"days_since_modify"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDaysSinceFileLastModifiedOperation) OpNum() int { return 37 }

func (o *xxx_SetDaysSinceFileLastModifiedOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/DaysSinceFileLastModified"
}

func (o *xxx_SetDaysSinceFileLastModifiedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDaysSinceFileLastModifiedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// daysSinceModify {in} (1:(int32))
	{
		if err := w.WriteData(o.DaysSinceModify); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDaysSinceFileLastModifiedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// daysSinceModify {in} (1:(int32))
	{
		if err := w.ReadData(&o.DaysSinceModify); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDaysSinceFileLastModifiedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDaysSinceFileLastModifiedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetDaysSinceFileLastModifiedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetDaysSinceFileLastModifiedRequest structure represents the DaysSinceFileLastModified operation request
type SetDaysSinceFileLastModifiedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	DaysSinceModify int32          `idl:"name:daysSinceModify" json:"days_since_modify"`
}

func (o *SetDaysSinceFileLastModifiedRequest) xxx_ToOp(ctx context.Context, op *xxx_SetDaysSinceFileLastModifiedOperation) *xxx_SetDaysSinceFileLastModifiedOperation {
	if op == nil {
		op = &xxx_SetDaysSinceFileLastModifiedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DaysSinceModify = o.DaysSinceModify
	return op
}

func (o *SetDaysSinceFileLastModifiedRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDaysSinceFileLastModifiedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DaysSinceModify = op.DaysSinceModify
}
func (o *SetDaysSinceFileLastModifiedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetDaysSinceFileLastModifiedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDaysSinceFileLastModifiedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDaysSinceFileLastModifiedResponse structure represents the DaysSinceFileLastModified operation response
type SetDaysSinceFileLastModifiedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DaysSinceFileLastModified return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDaysSinceFileLastModifiedResponse) xxx_ToOp(ctx context.Context, op *xxx_SetDaysSinceFileLastModifiedOperation) *xxx_SetDaysSinceFileLastModifiedOperation {
	if op == nil {
		op = &xxx_SetDaysSinceFileLastModifiedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetDaysSinceFileLastModifiedResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDaysSinceFileLastModifiedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDaysSinceFileLastModifiedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetDaysSinceFileLastModifiedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDaysSinceFileLastModifiedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertyConditionsOperation structure represents the PropertyConditions operation
type xxx_GetPropertyConditionsOperation struct {
	This               *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat   `idl:"name:That" json:"that"`
	PropertyConditions *fsrm.Collection `idl:"name:propertyConditions" json:"property_conditions"`
	Return             int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertyConditionsOperation) OpNum() int { return 38 }

func (o *xxx_GetPropertyConditionsOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/PropertyConditions"
}

func (o *xxx_GetPropertyConditionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyConditionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPropertyConditionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPropertyConditionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyConditionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// propertyConditions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		if o.PropertyConditions != nil {
			_ptr_propertyConditions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyConditions != nil {
					if err := o.PropertyConditions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Collection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyConditions, _ptr_propertyConditions); err != nil {
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

func (o *xxx_GetPropertyConditionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// propertyConditions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		_ptr_propertyConditions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyConditions == nil {
				o.PropertyConditions = &fsrm.Collection{}
			}
			if err := o.PropertyConditions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertyConditions := func(ptr interface{}) { o.PropertyConditions = *ptr.(**fsrm.Collection) }
		if err := w.ReadPointer(&o.PropertyConditions, _s_propertyConditions, _ptr_propertyConditions); err != nil {
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

// GetPropertyConditionsRequest structure represents the PropertyConditions operation request
type GetPropertyConditionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPropertyConditionsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPropertyConditionsOperation) *xxx_GetPropertyConditionsOperation {
	if op == nil {
		op = &xxx_GetPropertyConditionsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPropertyConditionsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertyConditionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertyConditionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertyConditionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertyConditionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertyConditionsResponse structure represents the PropertyConditions operation response
type GetPropertyConditionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// propertyConditions: Pointer to an IFsrmCollection interface pointer (section 3.2.4.2.1)
	// that upon completion contains pointers to the property conditions configured for
	// the file management job. A caller MUST release the collection received when the caller
	// is done with it.
	PropertyConditions *fsrm.Collection `idl:"name:propertyConditions" json:"property_conditions"`
	// Return: The PropertyConditions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertyConditionsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPropertyConditionsOperation) *xxx_GetPropertyConditionsOperation {
	if op == nil {
		op = &xxx_GetPropertyConditionsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PropertyConditions = o.PropertyConditions
	op.Return = o.Return
	return op
}

func (o *GetPropertyConditionsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertyConditionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertyConditions = op.PropertyConditions
	o.Return = op.Return
}
func (o *GetPropertyConditionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertyConditionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertyConditionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFromDateOperation structure represents the FromDate operation
type xxx_GetFromDateOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FromDate float64        `idl:"name:fromDate" json:"from_date"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFromDateOperation) OpNum() int { return 39 }

func (o *xxx_GetFromDateOperation) OpName() string { return "/IFsrmFileManagementJob/v1/FromDate" }

func (o *xxx_GetFromDateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFromDateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFromDateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFromDateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFromDateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fromDate {out, retval} (1:{pointer=ref}*(1))(2:{alias=DATE}(float64))
	{
		if err := w.WriteData(o.FromDate); err != nil {
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

func (o *xxx_GetFromDateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fromDate {out, retval} (1:{pointer=ref}*(1))(2:{alias=DATE}(float64))
	{
		if err := w.ReadData(&o.FromDate); err != nil {
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

// GetFromDateRequest structure represents the FromDate operation request
type GetFromDateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFromDateRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFromDateOperation) *xxx_GetFromDateOperation {
	if op == nil {
		op = &xxx_GetFromDateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetFromDateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFromDateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFromDateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFromDateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFromDateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFromDateResponse structure represents the FromDate operation response
type GetFromDateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FromDate float64        `idl:"name:fromDate" json:"from_date"`
	// Return: The FromDate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFromDateResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFromDateOperation) *xxx_GetFromDateOperation {
	if op == nil {
		op = &xxx_GetFromDateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.FromDate = o.FromDate
	op.Return = o.Return
	return op
}

func (o *GetFromDateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFromDateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FromDate = op.FromDate
	o.Return = op.Return
}
func (o *GetFromDateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFromDateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFromDateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFromDateOperation structure represents the FromDate operation
type xxx_SetFromDateOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FromDate float64        `idl:"name:fromDate" json:"from_date"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFromDateOperation) OpNum() int { return 40 }

func (o *xxx_SetFromDateOperation) OpName() string { return "/IFsrmFileManagementJob/v1/FromDate" }

func (o *xxx_SetFromDateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFromDateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fromDate {in} (1:{alias=DATE}(float64))
	{
		if err := w.WriteData(o.FromDate); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFromDateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fromDate {in} (1:{alias=DATE}(float64))
	{
		if err := w.ReadData(&o.FromDate); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFromDateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFromDateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFromDateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFromDateRequest structure represents the FromDate operation request
type SetFromDateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	FromDate float64        `idl:"name:fromDate" json:"from_date"`
}

func (o *SetFromDateRequest) xxx_ToOp(ctx context.Context, op *xxx_SetFromDateOperation) *xxx_SetFromDateOperation {
	if op == nil {
		op = &xxx_SetFromDateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.FromDate = o.FromDate
	return op
}

func (o *SetFromDateRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFromDateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FromDate = op.FromDate
}
func (o *SetFromDateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetFromDateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFromDateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFromDateResponse structure represents the FromDate operation response
type SetFromDateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The FromDate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFromDateResponse) xxx_ToOp(ctx context.Context, op *xxx_SetFromDateOperation) *xxx_SetFromDateOperation {
	if op == nil {
		op = &xxx_SetFromDateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetFromDateResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFromDateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFromDateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetFromDateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFromDateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTaskOperation structure represents the Task operation
type xxx_GetTaskOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskName *oaut.String   `idl:"name:taskName" json:"task_name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTaskOperation) OpNum() int { return 41 }

func (o *xxx_GetTaskOperation) OpName() string { return "/IFsrmFileManagementJob/v1/Task" }

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
	// taskName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
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
	// taskName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
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

func (o *GetTaskRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTaskOperation) *xxx_GetTaskOperation {
	if op == nil {
		op = &xxx_GetTaskOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetTaskRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTaskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTaskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
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
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskName *oaut.String   `idl:"name:taskName" json:"task_name"`
	// Return: The Task return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTaskResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTaskOperation) *xxx_GetTaskOperation {
	if op == nil {
		op = &xxx_GetTaskOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.TaskName = o.TaskName
	op.Return = o.Return
	return op
}

func (o *GetTaskResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTaskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskName = op.TaskName
	o.Return = op.Return
}
func (o *GetTaskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
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
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskName *oaut.String   `idl:"name:taskName" json:"task_name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetTaskOperation) OpNum() int { return 42 }

func (o *xxx_SetTaskOperation) OpName() string { return "/IFsrmFileManagementJob/v1/Task" }

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
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	TaskName *oaut.String   `idl:"name:taskName" json:"task_name"`
}

func (o *SetTaskRequest) xxx_ToOp(ctx context.Context, op *xxx_SetTaskOperation) *xxx_SetTaskOperation {
	if op == nil {
		op = &xxx_SetTaskOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.TaskName = o.TaskName
	return op
}

func (o *SetTaskRequest) xxx_FromOp(ctx context.Context, op *xxx_SetTaskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.TaskName = op.TaskName
}
func (o *SetTaskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
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

func (o *SetTaskResponse) xxx_ToOp(ctx context.Context, op *xxx_SetTaskOperation) *xxx_SetTaskOperation {
	if op == nil {
		op = &xxx_SetTaskOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetTaskResponse) xxx_FromOp(ctx context.Context, op *xxx_SetTaskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetTaskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetTaskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTaskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetParametersOperation structure represents the Parameters operation
type xxx_GetParametersOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Parameters *oaut.SafeArray `idl:"name:parameters" json:"parameters"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetParametersOperation) OpNum() int { return 43 }

func (o *xxx_GetParametersOperation) OpName() string { return "/IFsrmFileManagementJob/v1/Parameters" }

func (o *xxx_GetParametersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetParametersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetParametersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetParametersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetParametersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// parameters {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Parameters != nil {
			_ptr_parameters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Parameters != nil {
					if err := o.Parameters.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Parameters, _ptr_parameters); err != nil {
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

func (o *xxx_GetParametersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// parameters {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_parameters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Parameters == nil {
				o.Parameters = &oaut.SafeArray{}
			}
			if err := o.Parameters.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_parameters := func(ptr interface{}) { o.Parameters = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Parameters, _s_parameters, _ptr_parameters); err != nil {
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

// GetParametersRequest structure represents the Parameters operation request
type GetParametersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetParametersRequest) xxx_ToOp(ctx context.Context, op *xxx_GetParametersOperation) *xxx_GetParametersOperation {
	if op == nil {
		op = &xxx_GetParametersOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetParametersRequest) xxx_FromOp(ctx context.Context, op *xxx_GetParametersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetParametersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetParametersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetParametersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetParametersResponse structure represents the Parameters operation response
type GetParametersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Parameters *oaut.SafeArray `idl:"name:parameters" json:"parameters"`
	// Return: The Parameters return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetParametersResponse) xxx_ToOp(ctx context.Context, op *xxx_GetParametersOperation) *xxx_GetParametersOperation {
	if op == nil {
		op = &xxx_GetParametersOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Parameters = o.Parameters
	op.Return = o.Return
	return op
}

func (o *GetParametersResponse) xxx_FromOp(ctx context.Context, op *xxx_GetParametersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Parameters = op.Parameters
	o.Return = op.Return
}
func (o *GetParametersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetParametersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetParametersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetParametersOperation structure represents the Parameters operation
type xxx_SetParametersOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Parameters *oaut.SafeArray `idl:"name:parameters" json:"parameters"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetParametersOperation) OpNum() int { return 44 }

func (o *xxx_SetParametersOperation) OpName() string { return "/IFsrmFileManagementJob/v1/Parameters" }

func (o *xxx_SetParametersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetParametersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// parameters {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Parameters != nil {
			_ptr_parameters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Parameters != nil {
					if err := o.Parameters.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Parameters, _ptr_parameters); err != nil {
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

func (o *xxx_SetParametersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// parameters {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_parameters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Parameters == nil {
				o.Parameters = &oaut.SafeArray{}
			}
			if err := o.Parameters.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_parameters := func(ptr interface{}) { o.Parameters = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Parameters, _s_parameters, _ptr_parameters); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetParametersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetParametersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetParametersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetParametersRequest structure represents the Parameters operation request
type SetParametersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	Parameters *oaut.SafeArray `idl:"name:parameters" json:"parameters"`
}

func (o *SetParametersRequest) xxx_ToOp(ctx context.Context, op *xxx_SetParametersOperation) *xxx_SetParametersOperation {
	if op == nil {
		op = &xxx_SetParametersOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Parameters = o.Parameters
	return op
}

func (o *SetParametersRequest) xxx_FromOp(ctx context.Context, op *xxx_SetParametersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Parameters = op.Parameters
}
func (o *SetParametersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetParametersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetParametersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetParametersResponse structure represents the Parameters operation response
type SetParametersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Parameters return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetParametersResponse) xxx_ToOp(ctx context.Context, op *xxx_SetParametersOperation) *xxx_SetParametersOperation {
	if op == nil {
		op = &xxx_SetParametersOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetParametersResponse) xxx_FromOp(ctx context.Context, op *xxx_SetParametersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetParametersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetParametersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetParametersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRunningStatusOperation structure represents the RunningStatus operation
type xxx_GetRunningStatusOperation struct {
	This          *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat           `idl:"name:That" json:"that"`
	RunningStatus fsrm.ReportRunningStatus `idl:"name:runningStatus" json:"running_status"`
	Return        int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRunningStatusOperation) OpNum() int { return 45 }

func (o *xxx_GetRunningStatusOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/RunningStatus"
}

func (o *xxx_GetRunningStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRunningStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetRunningStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetRunningStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRunningStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// runningStatus {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmReportRunningStatus}(enum))
	{
		if err := w.WriteEnum(uint16(o.RunningStatus)); err != nil {
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

func (o *xxx_GetRunningStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// runningStatus {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmReportRunningStatus}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.RunningStatus)); err != nil {
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

// GetRunningStatusRequest structure represents the RunningStatus operation request
type GetRunningStatusRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetRunningStatusRequest) xxx_ToOp(ctx context.Context, op *xxx_GetRunningStatusOperation) *xxx_GetRunningStatusOperation {
	if op == nil {
		op = &xxx_GetRunningStatusOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetRunningStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRunningStatusOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetRunningStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetRunningStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRunningStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRunningStatusResponse structure represents the RunningStatus operation response
type GetRunningStatusResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat           `idl:"name:That" json:"that"`
	RunningStatus fsrm.ReportRunningStatus `idl:"name:runningStatus" json:"running_status"`
	// Return: The RunningStatus return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRunningStatusResponse) xxx_ToOp(ctx context.Context, op *xxx_GetRunningStatusOperation) *xxx_GetRunningStatusOperation {
	if op == nil {
		op = &xxx_GetRunningStatusOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.RunningStatus = o.RunningStatus
	op.Return = o.Return
	return op
}

func (o *GetRunningStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRunningStatusOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RunningStatus = op.RunningStatus
	o.Return = op.Return
}
func (o *GetRunningStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetRunningStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRunningStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLastErrorOperation structure represents the LastError operation
type xxx_GetLastErrorOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	LastError *oaut.String   `idl:"name:lastError" json:"last_error"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLastErrorOperation) OpNum() int { return 46 }

func (o *xxx_GetLastErrorOperation) OpName() string { return "/IFsrmFileManagementJob/v1/LastError" }

func (o *xxx_GetLastErrorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastErrorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLastErrorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLastErrorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastErrorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lastError {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.LastError != nil {
			_ptr_lastError := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.LastError != nil {
					if err := o.LastError.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LastError, _ptr_lastError); err != nil {
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

func (o *xxx_GetLastErrorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lastError {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_lastError := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.LastError == nil {
				o.LastError = &oaut.String{}
			}
			if err := o.LastError.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lastError := func(ptr interface{}) { o.LastError = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.LastError, _s_lastError, _ptr_lastError); err != nil {
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

// GetLastErrorRequest structure represents the LastError operation request
type GetLastErrorRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLastErrorRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLastErrorOperation) *xxx_GetLastErrorOperation {
	if op == nil {
		op = &xxx_GetLastErrorOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetLastErrorRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLastErrorOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLastErrorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLastErrorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastErrorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLastErrorResponse structure represents the LastError operation response
type GetLastErrorResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	LastError *oaut.String   `idl:"name:lastError" json:"last_error"`
	// Return: The LastError return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLastErrorResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLastErrorOperation) *xxx_GetLastErrorOperation {
	if op == nil {
		op = &xxx_GetLastErrorOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.LastError = o.LastError
	op.Return = o.Return
	return op
}

func (o *GetLastErrorResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLastErrorOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LastError = op.LastError
	o.Return = op.Return
}
func (o *GetLastErrorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLastErrorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastErrorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLastReportPathWithoutExtensionOperation structure represents the LastReportPathWithoutExtension operation
type xxx_GetLastReportPathWithoutExtensionOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   *oaut.String   `idl:"name:path" json:"path"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLastReportPathWithoutExtensionOperation) OpNum() int { return 47 }

func (o *xxx_GetLastReportPathWithoutExtensionOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/LastReportPathWithoutExtension"
}

func (o *xxx_GetLastReportPathWithoutExtensionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastReportPathWithoutExtensionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLastReportPathWithoutExtensionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLastReportPathWithoutExtensionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastReportPathWithoutExtensionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLastReportPathWithoutExtensionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetLastReportPathWithoutExtensionRequest structure represents the LastReportPathWithoutExtension operation request
type GetLastReportPathWithoutExtensionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLastReportPathWithoutExtensionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLastReportPathWithoutExtensionOperation) *xxx_GetLastReportPathWithoutExtensionOperation {
	if op == nil {
		op = &xxx_GetLastReportPathWithoutExtensionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetLastReportPathWithoutExtensionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLastReportPathWithoutExtensionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLastReportPathWithoutExtensionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLastReportPathWithoutExtensionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastReportPathWithoutExtensionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLastReportPathWithoutExtensionResponse structure represents the LastReportPathWithoutExtension operation response
type GetLastReportPathWithoutExtensionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// path: Pointer to a variable that upon completion contains the last report directory
	// without extension where the generated report(s) was (were) stored when the file management
	// job was previously run.
	Path *oaut.String `idl:"name:path" json:"path"`
	// Return: The LastReportPathWithoutExtension return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLastReportPathWithoutExtensionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLastReportPathWithoutExtensionOperation) *xxx_GetLastReportPathWithoutExtensionOperation {
	if op == nil {
		op = &xxx_GetLastReportPathWithoutExtensionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Path = o.Path
	op.Return = o.Return
	return op
}

func (o *GetLastReportPathWithoutExtensionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLastReportPathWithoutExtensionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Path = op.Path
	o.Return = op.Return
}
func (o *GetLastReportPathWithoutExtensionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLastReportPathWithoutExtensionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastReportPathWithoutExtensionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLastRunOperation structure represents the LastRun operation
type xxx_GetLastRunOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	LastRun float64        `idl:"name:lastRun" json:"last_run"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLastRunOperation) OpNum() int { return 48 }

func (o *xxx_GetLastRunOperation) OpName() string { return "/IFsrmFileManagementJob/v1/LastRun" }

func (o *xxx_GetLastRunOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastRunOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLastRunOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLastRunOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastRunOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lastRun {out, retval} (1:{pointer=ref}*(1))(2:{alias=DATE}(float64))
	{
		if err := w.WriteData(o.LastRun); err != nil {
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

func (o *xxx_GetLastRunOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lastRun {out, retval} (1:{pointer=ref}*(1))(2:{alias=DATE}(float64))
	{
		if err := w.ReadData(&o.LastRun); err != nil {
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

// GetLastRunRequest structure represents the LastRun operation request
type GetLastRunRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLastRunRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLastRunOperation) *xxx_GetLastRunOperation {
	if op == nil {
		op = &xxx_GetLastRunOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetLastRunRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLastRunOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLastRunRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLastRunRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastRunOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLastRunResponse structure represents the LastRun operation response
type GetLastRunResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	LastRun float64        `idl:"name:lastRun" json:"last_run"`
	// Return: The LastRun return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLastRunResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLastRunOperation) *xxx_GetLastRunOperation {
	if op == nil {
		op = &xxx_GetLastRunOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.LastRun = o.LastRun
	op.Return = o.Return
	return op
}

func (o *GetLastRunResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLastRunOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LastRun = op.LastRun
	o.Return = op.Return
}
func (o *GetLastRunResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLastRunResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastRunOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFileNamePatternOperation structure represents the FileNamePattern operation
type xxx_GetFileNamePatternOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileNamePattern *oaut.String   `idl:"name:fileNamePattern" json:"file_name_pattern"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFileNamePatternOperation) OpNum() int { return 49 }

func (o *xxx_GetFileNamePatternOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/FileNamePattern"
}

func (o *xxx_GetFileNamePatternOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileNamePatternOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFileNamePatternOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFileNamePatternOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileNamePatternOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileNamePattern {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FileNamePattern != nil {
			_ptr_fileNamePattern := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileNamePattern != nil {
					if err := o.FileNamePattern.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileNamePattern, _ptr_fileNamePattern); err != nil {
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

func (o *xxx_GetFileNamePatternOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileNamePattern {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_fileNamePattern := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileNamePattern == nil {
				o.FileNamePattern = &oaut.String{}
			}
			if err := o.FileNamePattern.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileNamePattern := func(ptr interface{}) { o.FileNamePattern = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FileNamePattern, _s_fileNamePattern, _ptr_fileNamePattern); err != nil {
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

// GetFileNamePatternRequest structure represents the FileNamePattern operation request
type GetFileNamePatternRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFileNamePatternRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFileNamePatternOperation) *xxx_GetFileNamePatternOperation {
	if op == nil {
		op = &xxx_GetFileNamePatternOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetFileNamePatternRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFileNamePatternOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFileNamePatternRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFileNamePatternRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileNamePatternOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFileNamePatternResponse structure represents the FileNamePattern operation response
type GetFileNamePatternResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileNamePattern *oaut.String   `idl:"name:fileNamePattern" json:"file_name_pattern"`
	// Return: The FileNamePattern return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFileNamePatternResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFileNamePatternOperation) *xxx_GetFileNamePatternOperation {
	if op == nil {
		op = &xxx_GetFileNamePatternOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.FileNamePattern = o.FileNamePattern
	op.Return = o.Return
	return op
}

func (o *GetFileNamePatternResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFileNamePatternOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileNamePattern = op.FileNamePattern
	o.Return = op.Return
}
func (o *GetFileNamePatternResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFileNamePatternResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileNamePatternOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFileNamePatternOperation structure represents the FileNamePattern operation
type xxx_SetFileNamePatternOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileNamePattern *oaut.String   `idl:"name:fileNamePattern" json:"file_name_pattern"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFileNamePatternOperation) OpNum() int { return 50 }

func (o *xxx_SetFileNamePatternOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/FileNamePattern"
}

func (o *xxx_SetFileNamePatternOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileNamePatternOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fileNamePattern {in} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FileNamePattern != nil {
			_ptr_fileNamePattern := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileNamePattern != nil {
					if err := o.FileNamePattern.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileNamePattern, _ptr_fileNamePattern); err != nil {
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

func (o *xxx_SetFileNamePatternOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fileNamePattern {in} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_fileNamePattern := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileNamePattern == nil {
				o.FileNamePattern = &oaut.String{}
			}
			if err := o.FileNamePattern.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileNamePattern := func(ptr interface{}) { o.FileNamePattern = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FileNamePattern, _s_fileNamePattern, _ptr_fileNamePattern); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileNamePatternOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileNamePatternOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFileNamePatternOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFileNamePatternRequest structure represents the FileNamePattern operation request
type SetFileNamePatternRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	FileNamePattern *oaut.String   `idl:"name:fileNamePattern" json:"file_name_pattern"`
}

func (o *SetFileNamePatternRequest) xxx_ToOp(ctx context.Context, op *xxx_SetFileNamePatternOperation) *xxx_SetFileNamePatternOperation {
	if op == nil {
		op = &xxx_SetFileNamePatternOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.FileNamePattern = o.FileNamePattern
	return op
}

func (o *SetFileNamePatternRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFileNamePatternOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FileNamePattern = op.FileNamePattern
}
func (o *SetFileNamePatternRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetFileNamePatternRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFileNamePatternOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFileNamePatternResponse structure represents the FileNamePattern operation response
type SetFileNamePatternResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The FileNamePattern return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFileNamePatternResponse) xxx_ToOp(ctx context.Context, op *xxx_SetFileNamePatternOperation) *xxx_SetFileNamePatternOperation {
	if op == nil {
		op = &xxx_SetFileNamePatternOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetFileNamePatternResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFileNamePatternOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFileNamePatternResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetFileNamePatternResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFileNamePatternOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RunOperation structure represents the Run operation
type xxx_RunOperation struct {
	This    *dcom.ORPCThis               `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat               `idl:"name:That" json:"that"`
	Context fsrm.ReportGenerationContext `idl:"name:context" json:"context"`
	Return  int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_RunOperation) OpNum() int { return 51 }

func (o *xxx_RunOperation) OpName() string { return "/IFsrmFileManagementJob/v1/Run" }

func (o *xxx_RunOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RunOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_RunOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RunOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RunRequest structure represents the Run operation request
type RunRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis               `idl:"name:This" json:"this"`
	Context fsrm.ReportGenerationContext `idl:"name:context" json:"context"`
}

func (o *RunRequest) xxx_ToOp(ctx context.Context, op *xxx_RunOperation) *xxx_RunOperation {
	if op == nil {
		op = &xxx_RunOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Context = o.Context
	return op
}

func (o *RunRequest) xxx_FromOp(ctx context.Context, op *xxx_RunOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Context = op.Context
}
func (o *RunRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RunRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RunOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RunResponse structure represents the Run operation response
type RunResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Run return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RunResponse) xxx_ToOp(ctx context.Context, op *xxx_RunOperation) *xxx_RunOperation {
	if op == nil {
		op = &xxx_RunOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RunResponse) xxx_FromOp(ctx context.Context, op *xxx_RunOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RunResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RunResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RunOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WaitForCompletionOperation structure represents the WaitForCompletion operation
type xxx_WaitForCompletionOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	WaitSeconds int32          `idl:"name:waitSeconds" json:"wait_seconds"`
	Completed   int16          `idl:"name:completed" json:"completed"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_WaitForCompletionOperation) OpNum() int { return 52 }

func (o *xxx_WaitForCompletionOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/WaitForCompletion"
}

func (o *xxx_WaitForCompletionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForCompletionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// waitSeconds {in} (1:(int32))
	{
		if err := w.WriteData(o.WaitSeconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForCompletionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// waitSeconds {in} (1:(int32))
	{
		if err := w.ReadData(&o.WaitSeconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForCompletionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForCompletionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// completed {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Completed); err != nil {
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

func (o *xxx_WaitForCompletionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// completed {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Completed); err != nil {
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

// WaitForCompletionRequest structure represents the WaitForCompletion operation request
type WaitForCompletionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	WaitSeconds int32          `idl:"name:waitSeconds" json:"wait_seconds"`
}

func (o *WaitForCompletionRequest) xxx_ToOp(ctx context.Context, op *xxx_WaitForCompletionOperation) *xxx_WaitForCompletionOperation {
	if op == nil {
		op = &xxx_WaitForCompletionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.WaitSeconds = o.WaitSeconds
	return op
}

func (o *WaitForCompletionRequest) xxx_FromOp(ctx context.Context, op *xxx_WaitForCompletionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WaitSeconds = op.WaitSeconds
}
func (o *WaitForCompletionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WaitForCompletionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitForCompletionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WaitForCompletionResponse structure represents the WaitForCompletion operation response
type WaitForCompletionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Completed int16          `idl:"name:completed" json:"completed"`
	// Return: The WaitForCompletion return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WaitForCompletionResponse) xxx_ToOp(ctx context.Context, op *xxx_WaitForCompletionOperation) *xxx_WaitForCompletionOperation {
	if op == nil {
		op = &xxx_WaitForCompletionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Completed = o.Completed
	op.Return = o.Return
	return op
}

func (o *WaitForCompletionResponse) xxx_FromOp(ctx context.Context, op *xxx_WaitForCompletionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Completed = op.Completed
	o.Return = op.Return
}
func (o *WaitForCompletionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WaitForCompletionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitForCompletionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CancelOperation structure represents the Cancel operation
type xxx_CancelOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CancelOperation) OpNum() int { return 53 }

func (o *xxx_CancelOperation) OpName() string { return "/IFsrmFileManagementJob/v1/Cancel" }

func (o *xxx_CancelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CancelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CancelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CancelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CancelRequest structure represents the Cancel operation request
type CancelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CancelRequest) xxx_ToOp(ctx context.Context, op *xxx_CancelOperation) *xxx_CancelOperation {
	if op == nil {
		op = &xxx_CancelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CancelRequest) xxx_FromOp(ctx context.Context, op *xxx_CancelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CancelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CancelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CancelResponse structure represents the Cancel operation response
type CancelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Cancel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CancelResponse) xxx_ToOp(ctx context.Context, op *xxx_CancelOperation) *xxx_CancelOperation {
	if op == nil {
		op = &xxx_CancelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CancelResponse) xxx_FromOp(ctx context.Context, op *xxx_CancelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CancelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CancelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddNotificationOperation structure represents the AddNotification operation
type xxx_AddNotificationOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Days   int32          `idl:"name:days" json:"days"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AddNotificationOperation) OpNum() int { return 54 }

func (o *xxx_AddNotificationOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/AddNotification"
}

func (o *xxx_AddNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// days {in} (1:(int32))
	{
		if err := w.WriteData(o.Days); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// days {in} (1:(int32))
	{
		if err := w.ReadData(&o.Days); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AddNotificationRequest structure represents the AddNotification operation request
type AddNotificationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// days: Contains the value of the notification period to add to the file management
	// job.
	Days int32 `idl:"name:days" json:"days"`
}

func (o *AddNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_AddNotificationOperation) *xxx_AddNotificationOperation {
	if op == nil {
		op = &xxx_AddNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Days = o.Days
	return op
}

func (o *AddNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_AddNotificationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Days = op.Days
}
func (o *AddNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddNotificationResponse structure represents the AddNotification operation response
type AddNotificationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AddNotification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_AddNotificationOperation) *xxx_AddNotificationOperation {
	if op == nil {
		op = &xxx_AddNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *AddNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_AddNotificationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AddNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteNotificationOperation structure represents the DeleteNotification operation
type xxx_DeleteNotificationOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Days   int32          `idl:"name:days" json:"days"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteNotificationOperation) OpNum() int { return 55 }

func (o *xxx_DeleteNotificationOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/DeleteNotification"
}

func (o *xxx_DeleteNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// days {in} (1:(int32))
	{
		if err := w.WriteData(o.Days); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// days {in} (1:(int32))
	{
		if err := w.ReadData(&o.Days); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteNotificationRequest structure represents the DeleteNotification operation request
type DeleteNotificationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// days: Contains the value of the notification period to delete from the file management
	// job.
	Days int32 `idl:"name:days" json:"days"`
}

func (o *DeleteNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteNotificationOperation) *xxx_DeleteNotificationOperation {
	if op == nil {
		op = &xxx_DeleteNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Days = o.Days
	return op
}

func (o *DeleteNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteNotificationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Days = op.Days
}
func (o *DeleteNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteNotificationResponse structure represents the DeleteNotification operation response
type DeleteNotificationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteNotification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteNotificationOperation) *xxx_DeleteNotificationOperation {
	if op == nil {
		op = &xxx_DeleteNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteNotificationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyNotificationOperation structure represents the ModifyNotification operation
type xxx_ModifyNotificationOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Days    int32          `idl:"name:days" json:"days"`
	NewDays int32          `idl:"name:newDays" json:"new_days"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyNotificationOperation) OpNum() int { return 56 }

func (o *xxx_ModifyNotificationOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/ModifyNotification"
}

func (o *xxx_ModifyNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// days {in} (1:(int32))
	{
		if err := w.WriteData(o.Days); err != nil {
			return err
		}
	}
	// newDays {in} (1:(int32))
	{
		if err := w.WriteData(o.NewDays); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// days {in} (1:(int32))
	{
		if err := w.ReadData(&o.Days); err != nil {
			return err
		}
	}
	// newDays {in} (1:(int32))
	{
		if err := w.ReadData(&o.NewDays); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ModifyNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ModifyNotificationRequest structure represents the ModifyNotification operation request
type ModifyNotificationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// days: Contains the value of the notification period to modify.
	Days int32 `idl:"name:days" json:"days"`
	// newDays: Contains the new value of the notification period.
	NewDays int32 `idl:"name:newDays" json:"new_days"`
}

func (o *ModifyNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_ModifyNotificationOperation) *xxx_ModifyNotificationOperation {
	if op == nil {
		op = &xxx_ModifyNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Days = o.Days
	op.NewDays = o.NewDays
	return op
}

func (o *ModifyNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifyNotificationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Days = op.Days
	o.NewDays = op.NewDays
}
func (o *ModifyNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ModifyNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyNotificationResponse structure represents the ModifyNotification operation response
type ModifyNotificationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ModifyNotification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ModifyNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_ModifyNotificationOperation) *xxx_ModifyNotificationOperation {
	if op == nil {
		op = &xxx_ModifyNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ModifyNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifyNotificationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ModifyNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ModifyNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateNotificationActionOperation structure represents the CreateNotificationAction operation
type xxx_CreateNotificationActionOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Days       int32           `idl:"name:days" json:"days"`
	ActionType fsrm.ActionType `idl:"name:actionType" json:"action_type"`
	Action     *fsrm.Action    `idl:"name:action" json:"action"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateNotificationActionOperation) OpNum() int { return 57 }

func (o *xxx_CreateNotificationActionOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/CreateNotificationAction"
}

func (o *xxx_CreateNotificationActionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNotificationActionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// days {in} (1:(int32))
	{
		if err := w.WriteData(o.Days); err != nil {
			return err
		}
	}
	// actionType {in} (1:{alias=FsrmActionType}(enum))
	{
		if err := w.WriteEnum(uint16(o.ActionType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNotificationActionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// days {in} (1:(int32))
	{
		if err := w.ReadData(&o.Days); err != nil {
			return err
		}
	}
	// actionType {in} (1:{alias=FsrmActionType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ActionType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNotificationActionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNotificationActionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// action {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmAction}(interface))
	{
		if o.Action != nil {
			_ptr_action := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Action != nil {
					if err := o.Action.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Action{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Action, _ptr_action); err != nil {
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

func (o *xxx_CreateNotificationActionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// action {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmAction}(interface))
	{
		_ptr_action := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Action == nil {
				o.Action = &fsrm.Action{}
			}
			if err := o.Action.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_action := func(ptr interface{}) { o.Action = *ptr.(**fsrm.Action) }
		if err := w.ReadPointer(&o.Action, _s_action, _ptr_action); err != nil {
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

// CreateNotificationActionRequest structure represents the CreateNotificationAction operation request
type CreateNotificationActionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// days: The days parameter contains the notification period for which to create the
	// action.
	Days int32 `idl:"name:days" json:"days"`
	// actionType: The actionType parameter contains the type of notification being created.
	ActionType fsrm.ActionType `idl:"name:actionType" json:"action_type"`
}

func (o *CreateNotificationActionRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateNotificationActionOperation) *xxx_CreateNotificationActionOperation {
	if op == nil {
		op = &xxx_CreateNotificationActionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Days = o.Days
	op.ActionType = o.ActionType
	return op
}

func (o *CreateNotificationActionRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateNotificationActionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Days = op.Days
	o.ActionType = op.ActionType
}
func (o *CreateNotificationActionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateNotificationActionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNotificationActionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateNotificationActionResponse structure represents the CreateNotificationAction operation response
type CreateNotificationActionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// action: Pointer to an IFsrmAction interface pointer (section 3.2.4.2.4) that upon
	// completion points to the newly created action. A caller MUST release the SAFEARRAY
	// received when the caller is done with it.
	Action *fsrm.Action `idl:"name:action" json:"action"`
	// Return: The CreateNotificationAction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateNotificationActionResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateNotificationActionOperation) *xxx_CreateNotificationActionOperation {
	if op == nil {
		op = &xxx_CreateNotificationActionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Action = o.Action
	op.Return = o.Return
	return op
}

func (o *CreateNotificationActionResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateNotificationActionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Action = op.Action
	o.Return = op.Return
}
func (o *CreateNotificationActionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateNotificationActionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNotificationActionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumNotificationActionsOperation structure represents the EnumNotificationActions operation
type xxx_EnumNotificationActionsOperation struct {
	This    *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Days    int32            `idl:"name:days" json:"days"`
	Actions *fsrm.Collection `idl:"name:actions" json:"actions"`
	Return  int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumNotificationActionsOperation) OpNum() int { return 58 }

func (o *xxx_EnumNotificationActionsOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/EnumNotificationActions"
}

func (o *xxx_EnumNotificationActionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumNotificationActionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// days {in} (1:(int32))
	{
		if err := w.WriteData(o.Days); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumNotificationActionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// days {in} (1:(int32))
	{
		if err := w.ReadData(&o.Days); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumNotificationActionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumNotificationActionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// actions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		if o.Actions != nil {
			_ptr_actions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Actions != nil {
					if err := o.Actions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Collection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Actions, _ptr_actions); err != nil {
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

func (o *xxx_EnumNotificationActionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// actions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		_ptr_actions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Actions == nil {
				o.Actions = &fsrm.Collection{}
			}
			if err := o.Actions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_actions := func(ptr interface{}) { o.Actions = *ptr.(**fsrm.Collection) }
		if err := w.ReadPointer(&o.Actions, _s_actions, _ptr_actions); err != nil {
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

// EnumNotificationActionsRequest structure represents the EnumNotificationActions operation request
type EnumNotificationActionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// days: The days parameter contains the notification period for which notifications
	// are to be enumerated.
	Days int32 `idl:"name:days" json:"days"`
}

func (o *EnumNotificationActionsRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumNotificationActionsOperation) *xxx_EnumNotificationActionsOperation {
	if op == nil {
		op = &xxx_EnumNotificationActionsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Days = o.Days
	return op
}

func (o *EnumNotificationActionsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumNotificationActionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Days = op.Days
}
func (o *EnumNotificationActionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumNotificationActionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumNotificationActionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumNotificationActionsResponse structure represents the EnumNotificationActions operation response
type EnumNotificationActionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// actions: Pointer to an IFsrmCollection interface pointer (section 3.2.4.2.1) that
	// upon completion contains IFsrmAction interface pointers of all the notifications
	// for the specified notification period. A caller MUST release the collection received
	// when the caller is done with it. To get the specific action interface for the action,
	// the caller MUST call QueryInterface for the interface corresponding to the action's
	// action type.
	Actions *fsrm.Collection `idl:"name:actions" json:"actions"`
	// Return: The EnumNotificationActions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumNotificationActionsResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumNotificationActionsOperation) *xxx_EnumNotificationActionsOperation {
	if op == nil {
		op = &xxx_EnumNotificationActionsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Actions = o.Actions
	op.Return = o.Return
	return op
}

func (o *EnumNotificationActionsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumNotificationActionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Actions = op.Actions
	o.Return = op.Return
}
func (o *EnumNotificationActionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumNotificationActionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumNotificationActionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreatePropertyConditionOperation structure represents the CreatePropertyCondition operation
type xxx_CreatePropertyConditionOperation struct {
	This              *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Name              *oaut.String            `idl:"name:name" json:"name"`
	PropertyCondition *fsrm.PropertyCondition `idl:"name:propertyCondition" json:"property_condition"`
	Return            int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_CreatePropertyConditionOperation) OpNum() int { return 59 }

func (o *xxx_CreatePropertyConditionOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/CreatePropertyCondition"
}

func (o *xxx_CreatePropertyConditionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePropertyConditionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_CreatePropertyConditionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_CreatePropertyConditionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePropertyConditionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// propertyCondition {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmPropertyCondition}(interface))
	{
		if o.PropertyCondition != nil {
			_ptr_propertyCondition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyCondition != nil {
					if err := o.PropertyCondition.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.PropertyCondition{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyCondition, _ptr_propertyCondition); err != nil {
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

func (o *xxx_CreatePropertyConditionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// propertyCondition {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmPropertyCondition}(interface))
	{
		_ptr_propertyCondition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyCondition == nil {
				o.PropertyCondition = &fsrm.PropertyCondition{}
			}
			if err := o.PropertyCondition.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertyCondition := func(ptr interface{}) { o.PropertyCondition = *ptr.(**fsrm.PropertyCondition) }
		if err := w.ReadPointer(&o.PropertyCondition, _s_propertyCondition, _ptr_propertyCondition); err != nil {
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

// CreatePropertyConditionRequest structure represents the CreatePropertyCondition operation request
type CreatePropertyConditionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// name: Contains the name of the property condition to create.
	Name *oaut.String `idl:"name:name" json:"name"`
}

func (o *CreatePropertyConditionRequest) xxx_ToOp(ctx context.Context, op *xxx_CreatePropertyConditionOperation) *xxx_CreatePropertyConditionOperation {
	if op == nil {
		op = &xxx_CreatePropertyConditionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Name = o.Name
	return op
}

func (o *CreatePropertyConditionRequest) xxx_FromOp(ctx context.Context, op *xxx_CreatePropertyConditionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
}
func (o *CreatePropertyConditionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreatePropertyConditionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePropertyConditionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreatePropertyConditionResponse structure represents the CreatePropertyCondition operation response
type CreatePropertyConditionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// propertyCondition: Pointer to an IFsrmPropertyCondition interface pointer (section
	// 3.2.4.2.49) that upon completion points to the newly created property condition.
	// A caller MUST release the SAFEARRAY received when it is done with it.
	PropertyCondition *fsrm.PropertyCondition `idl:"name:propertyCondition" json:"property_condition"`
	// Return: The CreatePropertyCondition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreatePropertyConditionResponse) xxx_ToOp(ctx context.Context, op *xxx_CreatePropertyConditionOperation) *xxx_CreatePropertyConditionOperation {
	if op == nil {
		op = &xxx_CreatePropertyConditionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PropertyCondition = o.PropertyCondition
	op.Return = o.Return
	return op
}

func (o *CreatePropertyConditionResponse) xxx_FromOp(ctx context.Context, op *xxx_CreatePropertyConditionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertyCondition = op.PropertyCondition
	o.Return = op.Return
}
func (o *CreatePropertyConditionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreatePropertyConditionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePropertyConditionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateCustomActionOperation structure represents the CreateCustomAction operation
type xxx_CreateCustomActionOperation struct {
	This         *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat      `idl:"name:That" json:"that"`
	CustomAction *fsrm.ActionCommand `idl:"name:customAction" json:"custom_action"`
	Return       int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateCustomActionOperation) OpNum() int { return 60 }

func (o *xxx_CreateCustomActionOperation) OpName() string {
	return "/IFsrmFileManagementJob/v1/CreateCustomAction"
}

func (o *xxx_CreateCustomActionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateCustomActionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateCustomActionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CreateCustomActionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateCustomActionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// customAction {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmActionCommand}(interface))
	{
		if o.CustomAction != nil {
			_ptr_customAction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CustomAction != nil {
					if err := o.CustomAction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.ActionCommand{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CustomAction, _ptr_customAction); err != nil {
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

func (o *xxx_CreateCustomActionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// customAction {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmActionCommand}(interface))
	{
		_ptr_customAction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CustomAction == nil {
				o.CustomAction = &fsrm.ActionCommand{}
			}
			if err := o.CustomAction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_customAction := func(ptr interface{}) { o.CustomAction = *ptr.(**fsrm.ActionCommand) }
		if err := w.ReadPointer(&o.CustomAction, _s_customAction, _ptr_customAction); err != nil {
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

// CreateCustomActionRequest structure represents the CreateCustomAction operation request
type CreateCustomActionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CreateCustomActionRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateCustomActionOperation) *xxx_CreateCustomActionOperation {
	if op == nil {
		op = &xxx_CreateCustomActionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CreateCustomActionRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateCustomActionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CreateCustomActionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateCustomActionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateCustomActionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateCustomActionResponse structure represents the CreateCustomAction operation response
type CreateCustomActionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// customAction: Pointer to an IFsrmActionCommand interface pointer (section 3.2.4.2.9)
	// that upon completion points to the newly created notification. A caller MUST release
	// the IFsrmActionCommand received when the caller is done with it.
	CustomAction *fsrm.ActionCommand `idl:"name:customAction" json:"custom_action"`
	// Return: The CreateCustomAction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateCustomActionResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateCustomActionOperation) *xxx_CreateCustomActionOperation {
	if op == nil {
		op = &xxx_CreateCustomActionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.CustomAction = o.CustomAction
	op.Return = o.Return
	return op
}

func (o *CreateCustomActionResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateCustomActionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CustomAction = op.CustomAction
	o.Return = op.Return
}
func (o *CreateCustomActionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateCustomActionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateCustomActionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
