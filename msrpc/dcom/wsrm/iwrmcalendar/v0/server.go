package iwrmcalendar

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

// IWRMCalendar server interface.
type CalendarServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The GetCalendarInfo method gets information about one or all calendar events.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | Operation successful.              |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// Additional IWRMCalendar interface methods are specified in section 3.2.4.4.
	GetCalendarInfo(context.Context, *GetCalendarInfoRequest) (*GetCalendarInfoResponse, error)

	// The CreateCalendar method creates a new calendar event.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                  |
	//	|                VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                           | Operation successful.                                                            |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                   | One or more arguments are invalid.                                               |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER      | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                           | be processed.<57>                                                                |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0258 WRM_ERR_CAL_DUPLICATE_CALENDAR | A calendar event with the specified name already exists.                         |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF025B WRM_ERR_CAL_MAX_CAL_EXCEEDED   | The number of calendar events that exist at one time has exceeded an             |
	//	|                                           | implementation-defined limit.<58>                                                |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMCalendar interface methods are specified in section 3.2.4.4.
	CreateCalendar(context.Context, *CreateCalendarRequest) (*CreateCalendarResponse, error)

	// The ModifyCalendar method modifies an existing calendar event.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                              | Operation successful.                                                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                      | One or more arguments are invalid.                                               |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0069 WRM_ERR_OLD_INFORMATION           | The XML timestamp is out of date.                                                |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER         | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                              | be processed.<60>                                                                |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0260 WRM_ERR_CAL_INTERVAL_TOO_LONG     | The calendar recurrence interval value MUST be from 1 to 999.                    |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0263 WRM_ERR_CAL_INVALID_MONTH         | The month value MUST be from 1 to 12.                                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0264 WRM_ERR_CAL_INVALID_MONTHDAY      | The days of the month MUST have a value from 1 to 31.                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0265 WRM_ERR_CAL_INVALID_YEARDAY       | The day of the year cannot be zero and SHOULD NOT exceed the number of days in   |
	//	|                                              | the selected year.<61>                                                           |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0266 WRM_ERR_CAL_INVALID_WEEKNO        | The weeks of the year MUST have a value from 1 to 53.                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0267 WRM_ERR_CAL_INVALID_SETPOS        | The value of the monthly instance of a day of the week is invalid.               |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0269 WRM_ERR_CAL_SCHEDULE_NAME_INVALID | The specified schedule name does not exist.                                      |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0275 WRM_ERR_CAL_INVALID_DURATION      | The specified calendar event duration is invalid. Specify a duration using the   |
	//	|                                              | format dd:hh:mm.                                                                 |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMCalendar interface methods are specified in section 3.2.4.4.
	ModifyCalendar(context.Context, *ModifyCalendarRequest) (*ModifyCalendarResponse, error)

	// The DeleteCalendar method deletes a specified calendar event.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                         | Operation successful.                                                            |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                 | One or more arguments are invalid.                                               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE             | The specified name contains characters that are invalid. The name cannot         |
	//	|                                         | start with a hyphen ("-") and cannot contain spaces or any of the following      |
	//	|                                         | characters: \ / ? * | : < > " , ;.                                               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0259 WRM_ERR_CAL_UNKNOWN_CALENDAR | The specified calendar event does not exist.                                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMCalendar interface methods are specified in section 3.2.4.4.
	DeleteCalendar(context.Context, *DeleteCalendarRequest) (*DeleteCalendarResponse, error)

	// The RenameCalendar method renames a specified calendar event.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                  |
	//	|                VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                           | Operation successful.                                                            |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                   | One or more arguments are invalid.                                               |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE               | The specified name contains characters that are invalid. The name cannot         |
	//	|                                           | start with a hyphen ("-") and cannot contain spaces or any of the following      |
	//	|                                           | characters: \ / ? * | : < > " , ;.                                               |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0258 WRM_ERR_CAL_DUPLICATE_CALENDAR | A calendar event with the specified name already exists.                         |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0259 WRM_ERR_CAL_UNKNOWN_CALENDAR   | The specified calendar event does not exist.                                     |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF025A WRM_ERR_CAL_NAME_TOO_LONG      | The calendar name has exceeded an implementation-defined limit.<63>.             |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMCalendar interface methods are specified in section 3.2.4.4.
	RenameCalendar(context.Context, *RenameCalendarRequest) (*RenameCalendarResponse, error)

	// The ComputeEvents method computes the calendar events in a specified time interval.
	// It returns the list of all calendar events from the system in pbstrEvents with a
	// start time greater than or equal to the one specified in szStartTime and with an
	// end time less than or equal to the one specified in szEndTime. If the method is called
	// with fMergeEvents set to TRUE, conflicting events are merged and returned in pbstrConflicts.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | Operation successful.              |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// Additional IWRMCalendar interface methods are specified in section 3.2.4.4.
	ComputeEvents(context.Context, *ComputeEventsRequest) (*ComputeEventsResponse, error)

	// The GetScheduleInfo method gets information about a specified schedule object.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-----------------------------------------+-----------------------------------------------+
	//	|                 RETURN                  |                                               |
	//	|               VALUE/CODE                |                  DESCRIPTION                  |
	//	|                                         |                                               |
	//	+-----------------------------------------+-----------------------------------------------+
	//	+-----------------------------------------+-----------------------------------------------+
	//	| 0x00000000 S_OK                         | Operation successful.                         |
	//	+-----------------------------------------+-----------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                 | One or more arguments are invalid.            |
	//	+-----------------------------------------+-----------------------------------------------+
	//	| 0xC1FF0271 WRM_ERR_CAL_UNKNOWN_SCHEDULE | The specified schedule object does not exist. |
	//	+-----------------------------------------+-----------------------------------------------+
	//
	// Additional IWRMCalendar interface methods are specified in section 3.2.4.4.
	GetScheduleInfo(context.Context, *GetScheduleInfoRequest) (*GetScheduleInfoResponse, error)

	// The CreateSchedule method creates a new schedule object.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                  |
	//	|                VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                           | Operation successful.                                                            |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                   | One or more arguments are invalid.                                               |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER      | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                           | be processed.<64>                                                                |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF025C WRM_ERR_CAL_MAX_SCHED_EXCEEDED | The number of schedules has exceeded an implementation-defined limit.<65>        |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0270 WRM_ERR_CAL_DUPLICATE_SCHEDULE | The specified schedule object already exists.                                    |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMCalendar interface methods are specified in section 3.2.4.4.
	CreateSchedule(context.Context, *CreateScheduleRequest) (*CreateScheduleResponse, error)

	// The ModifySchedule method modifies the specified schedule of the calendar.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                         | Operation successful.                                                            |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                 | One or more arguments are invalid.                                               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0069 WRM_ERR_OLD_INFORMATION      | The XML timestamp is out of date.                                                |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER    | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                         | be processed.<67>                                                                |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0271 WRM_ERR_CAL_UNKNOWN_SCHEDULE | The specified schedule object does not exist.                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMCalendar interface methods are specified in section 3.2.4.4.
	ModifySchedule(context.Context, *ModifyScheduleRequest) (*ModifyScheduleResponse, error)

	// The DeleteSchedule method deletes an existing schedule object.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                         | Operation successful.                                                            |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0272 WRM_ERR_CAL_SCHEDULE_IN_USE  | The selected schedule is currently in use with a calendar. Delete or edit the    |
	//	|                                         | calendar first.                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0271 WRM_ERR_CAL_UNKNOWN_SCHEDULE | The specified schedule object does not exist.                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                 | One or more arguments are invalid.                                               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMCalendar interface methods are specified in section 3.2.4.4.
	DeleteSchedule(context.Context, *DeleteScheduleRequest) (*DeleteScheduleResponse, error)

	// The RenameSchedule method renames a specified schedule object. If the schedule object
	// is being referenced by some calendar object, then the calendar object is also updated
	// with the new name.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                              | Operation successful.                                                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                      | One or more arguments are invalid.                                               |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE                  | The specified name contains characters that are invalid. The name cannot         |
	//	|                                              | start with a hyphen ("-") and cannot contain spaces or any of the following      |
	//	|                                              | characters: \ / ? * | : < > " , ;.                                               |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0270 WRM_ERR_CAL_DUPLICATE_SCHEDULE    | The new schedule name is already taken by an existing schedule object.           |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0271 WRM_ERR_CAL_UNKNOWN_SCHEDULE      | The specified schedule object does not exist.                                    |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0273 WRM_ERR_CAL_SCHEDULE_NAME_TOOLONG | The schedule object name has exceeded an implementation-defined limit.<68>       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMCalendar interface methods are specified in section 3.2.4.4.
	RenameSchedule(context.Context, *RenameScheduleRequest) (*RenameScheduleResponse, error)

	// The MoveBeforeCalendar method moves a calendar event before the specified reference
	// event. The caller can choose to control the move depending on whether the current
	// resource policy is affected.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-----------------------------------------+----------------------------------------------+
	//	|                 RETURN                  |                                              |
	//	|               VALUE/CODE                |                 DESCRIPTION                  |
	//	|                                         |                                              |
	//	+-----------------------------------------+----------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------+
	//	| 0x00000000 S_OK                         | Operation successful.                        |
	//	+-----------------------------------------+----------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                 | One or more arguments are invalid.           |
	//	+-----------------------------------------+----------------------------------------------+
	//	| 0xC1FF0259 WRM_ERR_CAL_UNKNOWN_CALENDAR | The specified calendar event does not exist. |
	//	+-----------------------------------------+----------------------------------------------+
	//
	// Additional IWRMCalendar interface methods are specified in section 3.2.4.4.
	MoveBeforeCalendar(context.Context, *MoveBeforeCalendarRequest) (*MoveBeforeCalendarResponse, error)

	// The MoveAfterCalendar method moves a calendar event after the specified reference
	// event. The caller can choose to control the move depending on whether the current
	// resource policy is affected.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-----------------------------------------+----------------------------------------------+
	//	|                 RETURN                  |                                              |
	//	|               VALUE/CODE                |                 DESCRIPTION                  |
	//	|                                         |                                              |
	//	+-----------------------------------------+----------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------+
	//	| 0x00000000 S_OK                         | Operation successful.                        |
	//	+-----------------------------------------+----------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                 | One or more arguments are invalid.           |
	//	+-----------------------------------------+----------------------------------------------+
	//	| 0xC1FF0259 WRM_ERR_CAL_UNKNOWN_CALENDAR | The specified calendar event does not exist. |
	//	+-----------------------------------------+----------------------------------------------+
	//
	// Additional IWRMCalendar interface methods are specified in section 3.2.4.4.
	MoveAfterCalendar(context.Context, *MoveAfterCalendarRequest) (*MoveAfterCalendarResponse, error)

	// The GetServerTimeZone method gets the current time zone setting of the server.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | Operation successful.              |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// Additional IWRMCalendar interface methods are specified in section 3.2.4.4.
	GetServerTimeZone(context.Context, *GetServerTimeZoneRequest) (*GetServerTimeZoneResponse, error)
}

func RegisterCalendarServer(conn dcerpc.Conn, o CalendarServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCalendarServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CalendarSyntaxV0_0))...)
}

func NewCalendarServerHandle(o CalendarServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CalendarServerHandle(ctx, o, opNum, r)
	}
}

func CalendarServerHandle(ctx context.Context, o CalendarServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // GetCalendarInfo
		op := &xxx_GetCalendarInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCalendarInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCalendarInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // CreateCalendar
		op := &xxx_CreateCalendarOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateCalendarRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateCalendar(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ModifyCalendar
		op := &xxx_ModifyCalendarOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyCalendarRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyCalendar(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // DeleteCalendar
		op := &xxx_DeleteCalendarOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteCalendarRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteCalendar(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RenameCalendar
		op := &xxx_RenameCalendarOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RenameCalendarRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RenameCalendar(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // ComputeEvents
		op := &xxx_ComputeEventsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ComputeEventsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ComputeEvents(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // GetScheduleInfo
		op := &xxx_GetScheduleInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetScheduleInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetScheduleInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // CreateSchedule
		op := &xxx_CreateScheduleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateScheduleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateSchedule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // ModifySchedule
		op := &xxx_ModifyScheduleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyScheduleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifySchedule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // DeleteSchedule
		op := &xxx_DeleteScheduleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteScheduleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteSchedule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // RenameSchedule
		op := &xxx_RenameScheduleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RenameScheduleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RenameSchedule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // MoveBeforeCalendar
		op := &xxx_MoveBeforeCalendarOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MoveBeforeCalendarRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MoveBeforeCalendar(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // MoveAfterCalendar
		op := &xxx_MoveAfterCalendarOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MoveAfterCalendarRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MoveAfterCalendar(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // GetServerTimeZone
		op := &xxx_GetServerTimeZoneOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServerTimeZoneRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServerTimeZone(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWRMCalendar
type UnimplementedCalendarServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedCalendarServer) GetCalendarInfo(context.Context, *GetCalendarInfoRequest) (*GetCalendarInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCalendarServer) CreateCalendar(context.Context, *CreateCalendarRequest) (*CreateCalendarResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCalendarServer) ModifyCalendar(context.Context, *ModifyCalendarRequest) (*ModifyCalendarResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCalendarServer) DeleteCalendar(context.Context, *DeleteCalendarRequest) (*DeleteCalendarResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCalendarServer) RenameCalendar(context.Context, *RenameCalendarRequest) (*RenameCalendarResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCalendarServer) ComputeEvents(context.Context, *ComputeEventsRequest) (*ComputeEventsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCalendarServer) GetScheduleInfo(context.Context, *GetScheduleInfoRequest) (*GetScheduleInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCalendarServer) CreateSchedule(context.Context, *CreateScheduleRequest) (*CreateScheduleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCalendarServer) ModifySchedule(context.Context, *ModifyScheduleRequest) (*ModifyScheduleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCalendarServer) DeleteSchedule(context.Context, *DeleteScheduleRequest) (*DeleteScheduleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCalendarServer) RenameSchedule(context.Context, *RenameScheduleRequest) (*RenameScheduleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCalendarServer) MoveBeforeCalendar(context.Context, *MoveBeforeCalendarRequest) (*MoveBeforeCalendarResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCalendarServer) MoveAfterCalendar(context.Context, *MoveAfterCalendarRequest) (*MoveAfterCalendarResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCalendarServer) GetServerTimeZone(context.Context, *GetServerTimeZoneRequest) (*GetServerTimeZoneResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CalendarServer = (*UnimplementedCalendarServer)(nil)
