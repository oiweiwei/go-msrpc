package iwrmcalendar

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
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wsrm"
)

var (
	// IWRMCalendar interface identifier 481e06cf-ab04-4498-8ffe-124a0a34296d
	CalendarIID = &dcom.IID{Data1: 0x481e06cf, Data2: 0xab04, Data3: 0x4498, Data4: []byte{0x8f, 0xfe, 0x12, 0x4a, 0x0a, 0x34, 0x29, 0x6d}}
	// Syntax UUID
	CalendarSyntaxUUID = &uuid.UUID{TimeLow: 0x481e06cf, TimeMid: 0xab04, TimeHiAndVersion: 0x4498, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xfe, Node: [6]uint8{0x12, 0x4a, 0xa, 0x34, 0x29, 0x6d}}
	// Syntax ID
	CalendarSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CalendarSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWRMCalendar interface.
type CalendarClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

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
	GetCalendarInfo(context.Context, *GetCalendarInfoRequest, ...dcerpc.CallOption) (*GetCalendarInfoResponse, error)

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
	CreateCalendar(context.Context, *CreateCalendarRequest, ...dcerpc.CallOption) (*CreateCalendarResponse, error)

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
	ModifyCalendar(context.Context, *ModifyCalendarRequest, ...dcerpc.CallOption) (*ModifyCalendarResponse, error)

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
	DeleteCalendar(context.Context, *DeleteCalendarRequest, ...dcerpc.CallOption) (*DeleteCalendarResponse, error)

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
	RenameCalendar(context.Context, *RenameCalendarRequest, ...dcerpc.CallOption) (*RenameCalendarResponse, error)

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
	ComputeEvents(context.Context, *ComputeEventsRequest, ...dcerpc.CallOption) (*ComputeEventsResponse, error)

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
	GetScheduleInfo(context.Context, *GetScheduleInfoRequest, ...dcerpc.CallOption) (*GetScheduleInfoResponse, error)

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
	CreateSchedule(context.Context, *CreateScheduleRequest, ...dcerpc.CallOption) (*CreateScheduleResponse, error)

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
	ModifySchedule(context.Context, *ModifyScheduleRequest, ...dcerpc.CallOption) (*ModifyScheduleResponse, error)

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
	DeleteSchedule(context.Context, *DeleteScheduleRequest, ...dcerpc.CallOption) (*DeleteScheduleResponse, error)

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
	RenameSchedule(context.Context, *RenameScheduleRequest, ...dcerpc.CallOption) (*RenameScheduleResponse, error)

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
	MoveBeforeCalendar(context.Context, *MoveBeforeCalendarRequest, ...dcerpc.CallOption) (*MoveBeforeCalendarResponse, error)

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
	MoveAfterCalendar(context.Context, *MoveAfterCalendarRequest, ...dcerpc.CallOption) (*MoveAfterCalendarResponse, error)

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
	GetServerTimeZone(context.Context, *GetServerTimeZoneRequest, ...dcerpc.CallOption) (*GetServerTimeZoneResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CalendarClient
}

type xxx_DefaultCalendarClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCalendarClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultCalendarClient) GetCalendarInfo(ctx context.Context, in *GetCalendarInfoRequest, opts ...dcerpc.CallOption) (*GetCalendarInfoResponse, error) {
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
	out := &GetCalendarInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCalendarClient) CreateCalendar(ctx context.Context, in *CreateCalendarRequest, opts ...dcerpc.CallOption) (*CreateCalendarResponse, error) {
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
	out := &CreateCalendarResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCalendarClient) ModifyCalendar(ctx context.Context, in *ModifyCalendarRequest, opts ...dcerpc.CallOption) (*ModifyCalendarResponse, error) {
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
	out := &ModifyCalendarResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCalendarClient) DeleteCalendar(ctx context.Context, in *DeleteCalendarRequest, opts ...dcerpc.CallOption) (*DeleteCalendarResponse, error) {
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
	out := &DeleteCalendarResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCalendarClient) RenameCalendar(ctx context.Context, in *RenameCalendarRequest, opts ...dcerpc.CallOption) (*RenameCalendarResponse, error) {
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
	out := &RenameCalendarResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCalendarClient) ComputeEvents(ctx context.Context, in *ComputeEventsRequest, opts ...dcerpc.CallOption) (*ComputeEventsResponse, error) {
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
	out := &ComputeEventsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCalendarClient) GetScheduleInfo(ctx context.Context, in *GetScheduleInfoRequest, opts ...dcerpc.CallOption) (*GetScheduleInfoResponse, error) {
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
	out := &GetScheduleInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCalendarClient) CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...dcerpc.CallOption) (*CreateScheduleResponse, error) {
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
	out := &CreateScheduleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCalendarClient) ModifySchedule(ctx context.Context, in *ModifyScheduleRequest, opts ...dcerpc.CallOption) (*ModifyScheduleResponse, error) {
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
	out := &ModifyScheduleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCalendarClient) DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, opts ...dcerpc.CallOption) (*DeleteScheduleResponse, error) {
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
	out := &DeleteScheduleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCalendarClient) RenameSchedule(ctx context.Context, in *RenameScheduleRequest, opts ...dcerpc.CallOption) (*RenameScheduleResponse, error) {
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
	out := &RenameScheduleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCalendarClient) MoveBeforeCalendar(ctx context.Context, in *MoveBeforeCalendarRequest, opts ...dcerpc.CallOption) (*MoveBeforeCalendarResponse, error) {
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
	out := &MoveBeforeCalendarResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCalendarClient) MoveAfterCalendar(ctx context.Context, in *MoveAfterCalendarRequest, opts ...dcerpc.CallOption) (*MoveAfterCalendarResponse, error) {
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
	out := &MoveAfterCalendarResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCalendarClient) GetServerTimeZone(ctx context.Context, in *GetServerTimeZoneRequest, opts ...dcerpc.CallOption) (*GetServerTimeZoneResponse, error) {
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
	out := &GetServerTimeZoneResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCalendarClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCalendarClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultCalendarClient) IPID(ctx context.Context, ipid *dcom.IPID) CalendarClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCalendarClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewCalendarClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CalendarClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CalendarSyntaxV0_0))...)
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
	return &xxx_DefaultCalendarClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetCalendarInfoOperation structure represents the GetCalendarInfo operation
type xxx_GetCalendarInfoOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	CalendarName *oaut.String   `idl:"name:bstrCalendarName" json:"calendar_name"`
	CalendarXML  *oaut.String   `idl:"name:pbstrCalendarXML" json:"calendar_xml"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCalendarInfoOperation) OpNum() int { return 7 }

func (o *xxx_GetCalendarInfoOperation) OpName() string { return "/IWRMCalendar/v0/GetCalendarInfo" }

func (o *xxx_GetCalendarInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCalendarInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CalendarName != nil {
			_ptr_bstrCalendarName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CalendarName != nil {
					if err := o.CalendarName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CalendarName, _ptr_bstrCalendarName); err != nil {
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

func (o *xxx_GetCalendarInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrCalendarName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CalendarName == nil {
				o.CalendarName = &oaut.String{}
			}
			if err := o.CalendarName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrCalendarName := func(ptr interface{}) { o.CalendarName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CalendarName, _s_bstrCalendarName, _ptr_bstrCalendarName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCalendarInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCalendarInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrCalendarXML {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CalendarXML != nil {
			_ptr_pbstrCalendarXML := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CalendarXML != nil {
					if err := o.CalendarXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CalendarXML, _ptr_pbstrCalendarXML); err != nil {
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

func (o *xxx_GetCalendarInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrCalendarXML {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrCalendarXML := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CalendarXML == nil {
				o.CalendarXML = &oaut.String{}
			}
			if err := o.CalendarXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrCalendarXML := func(ptr interface{}) { o.CalendarXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CalendarXML, _s_pbstrCalendarXML, _ptr_pbstrCalendarXML); err != nil {
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

// GetCalendarInfoRequest structure represents the GetCalendarInfo operation request
type GetCalendarInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrCalendarName: A string that specifies the name of the calendar event for which
	// information will be returned. If the string is "\", all calendar events are specified.
	CalendarName *oaut.String `idl:"name:bstrCalendarName" json:"calendar_name"`
}

func (o *GetCalendarInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCalendarInfoOperation) *xxx_GetCalendarInfoOperation {
	if op == nil {
		op = &xxx_GetCalendarInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.CalendarName = o.CalendarName
	return op
}

func (o *GetCalendarInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCalendarInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CalendarName = op.CalendarName
}
func (o *GetCalendarInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCalendarInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCalendarInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCalendarInfoResponse structure represents the GetCalendarInfo operation response
type GetCalendarInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrCalendarXML: A pointer to a string that returns the specified calendar information,
	// in the form of a Calendar element (section 2.2.5.7). For an example, see the Calendar
	// example  (section 4.2.6).
	//
	// If bstrCalendarName is "\", all calendar events MUST be returned, in the form of
	// a Calendars element (section 2.2.5.10). For an example, see the Calendars example
	// (section 4.2.7).
	CalendarXML *oaut.String `idl:"name:pbstrCalendarXML" json:"calendar_xml"`
	// Return: The GetCalendarInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCalendarInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCalendarInfoOperation) *xxx_GetCalendarInfoOperation {
	if op == nil {
		op = &xxx_GetCalendarInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.CalendarXML = o.CalendarXML
	op.Return = o.Return
	return op
}

func (o *GetCalendarInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCalendarInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CalendarXML = op.CalendarXML
	o.Return = op.Return
}
func (o *GetCalendarInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCalendarInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCalendarInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateCalendarOperation structure represents the CreateCalendar operation
type xxx_CreateCalendarOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	CalendarXML        *oaut.String   `idl:"name:bstrCalendarXML" json:"calendar_xml"`
	ChangeActivePolicy bool           `idl:"name:bChangeActivePolicy" json:"change_active_policy"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateCalendarOperation) OpNum() int { return 8 }

func (o *xxx_CreateCalendarOperation) OpName() string { return "/IWRMCalendar/v0/CreateCalendar" }

func (o *xxx_CreateCalendarOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateCalendarOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrCalendarXML {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CalendarXML != nil {
			_ptr_bstrCalendarXML := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CalendarXML != nil {
					if err := o.CalendarXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CalendarXML, _ptr_bstrCalendarXML); err != nil {
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
	// bChangeActivePolicy {in} (1:{alias=BOOL}(int32))
	{
		if !o.ChangeActivePolicy {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreateCalendarOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrCalendarXML {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrCalendarXML := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CalendarXML == nil {
				o.CalendarXML = &oaut.String{}
			}
			if err := o.CalendarXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrCalendarXML := func(ptr interface{}) { o.CalendarXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CalendarXML, _s_bstrCalendarXML, _ptr_bstrCalendarXML); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bChangeActivePolicy {in} (1:{alias=BOOL}(int32))
	{
		var _bChangeActivePolicy int32
		if err := w.ReadData(&_bChangeActivePolicy); err != nil {
			return err
		}
		o.ChangeActivePolicy = _bChangeActivePolicy != 0
	}
	return nil
}

func (o *xxx_CreateCalendarOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateCalendarOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateCalendarOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateCalendarRequest structure represents the CreateCalendar operation request
type CreateCalendarRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrCalendarXML: A string that specifies the new calendar event, in the form of a
	// Calendar element (section 2.2.5.7). For an example, see Calendar Example (section
	// 4.2.8).
	CalendarXML *oaut.String `idl:"name:bstrCalendarXML" json:"calendar_xml"`
	// bChangeActivePolicy: A Boolean value that specifies whether the configuration changes
	// made by this method call SHOULD change the current active policy, if applicable,
	// of the system.<56>
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | The current active policy of the system SHOULD NOT be changed by the method      |
	//	|                  | call.                                                                            |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | If applicable, the current active policy of the system SHOULD be changed by the  |
	//	|                  | method call.                                                                     |
	//	+------------------+----------------------------------------------------------------------------------+
	ChangeActivePolicy bool `idl:"name:bChangeActivePolicy" json:"change_active_policy"`
}

func (o *CreateCalendarRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateCalendarOperation) *xxx_CreateCalendarOperation {
	if op == nil {
		op = &xxx_CreateCalendarOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.CalendarXML = o.CalendarXML
	op.ChangeActivePolicy = o.ChangeActivePolicy
	return op
}

func (o *CreateCalendarRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateCalendarOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CalendarXML = op.CalendarXML
	o.ChangeActivePolicy = op.ChangeActivePolicy
}
func (o *CreateCalendarRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateCalendarRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateCalendarOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateCalendarResponse structure represents the CreateCalendar operation response
type CreateCalendarResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateCalendar return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateCalendarResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateCalendarOperation) *xxx_CreateCalendarOperation {
	if op == nil {
		op = &xxx_CreateCalendarOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CreateCalendarResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateCalendarOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateCalendarResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateCalendarResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateCalendarOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyCalendarOperation structure represents the ModifyCalendar operation
type xxx_ModifyCalendarOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	CalendarXML        *oaut.String   `idl:"name:bstrCalendarXML" json:"calendar_xml"`
	Overwrite          bool           `idl:"name:bOverwrite" json:"overwrite"`
	ChangeActivePolicy bool           `idl:"name:bChangeActivePolicy" json:"change_active_policy"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyCalendarOperation) OpNum() int { return 9 }

func (o *xxx_ModifyCalendarOperation) OpName() string { return "/IWRMCalendar/v0/ModifyCalendar" }

func (o *xxx_ModifyCalendarOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyCalendarOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrCalendarXML {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CalendarXML != nil {
			_ptr_bstrCalendarXML := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CalendarXML != nil {
					if err := o.CalendarXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CalendarXML, _ptr_bstrCalendarXML); err != nil {
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
	// bOverwrite {in} (1:{alias=BOOL}(int32))
	{
		if !o.Overwrite {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// bChangeActivePolicy {in} (1:{alias=BOOL}(int32))
	{
		if !o.ChangeActivePolicy {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ModifyCalendarOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrCalendarXML {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrCalendarXML := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CalendarXML == nil {
				o.CalendarXML = &oaut.String{}
			}
			if err := o.CalendarXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrCalendarXML := func(ptr interface{}) { o.CalendarXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CalendarXML, _s_bstrCalendarXML, _ptr_bstrCalendarXML); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bOverwrite {in} (1:{alias=BOOL}(int32))
	{
		var _bOverwrite int32
		if err := w.ReadData(&_bOverwrite); err != nil {
			return err
		}
		o.Overwrite = _bOverwrite != 0
	}
	// bChangeActivePolicy {in} (1:{alias=BOOL}(int32))
	{
		var _bChangeActivePolicy int32
		if err := w.ReadData(&_bChangeActivePolicy); err != nil {
			return err
		}
		o.ChangeActivePolicy = _bChangeActivePolicy != 0
	}
	return nil
}

func (o *xxx_ModifyCalendarOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyCalendarOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ModifyCalendarOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ModifyCalendarRequest structure represents the ModifyCalendar operation request
type ModifyCalendarRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrCalendarXML: A string that specifies the calendar event, in the form of a Calendar
	// element (section 2.2.5.7).
	CalendarXML *oaut.String `idl:"name:bstrCalendarXML" json:"calendar_xml"`
	// bOverwrite: A Boolean value that specifies whether to ignore the timestamp of the
	// specified calendar event when validating.
	//
	// A timestamp MUST be defined inside a common node at the root level of an XML element,
	// as shown in the Calendar example (section 4.2.6). The format of a timestamp is specified
	// in section 2.2.1.4.
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | The timestamp of the new calendar event MUST specify a time that is later        |
	//	|                  | than or equal to the timestamp of any modifications made to a calendar or        |
	//	|                  | schedule object on the server. Otherwise, the modification SHOULD fail, and      |
	//	|                  | WRM_ERR_OLD_INFORMATION SHOULD be returned.                                      |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | The calendar event is validated and modified without checking the timestamp.     |
	//	+------------------+----------------------------------------------------------------------------------+
	Overwrite bool `idl:"name:bOverwrite" json:"overwrite"`
	// bChangeActivePolicy: A Boolean value that specifies whether the configuration changes
	// made by this method call SHOULD change the current active policy, if applicable,
	// of the system.<59>
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | The current active policy of the system SHOULD NOT be changed by the method      |
	//	|                  | call.                                                                            |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | If applicable, the current active policy of the system SHOULD be changed by the  |
	//	|                  | method call.                                                                     |
	//	+------------------+----------------------------------------------------------------------------------+
	ChangeActivePolicy bool `idl:"name:bChangeActivePolicy" json:"change_active_policy"`
}

func (o *ModifyCalendarRequest) xxx_ToOp(ctx context.Context, op *xxx_ModifyCalendarOperation) *xxx_ModifyCalendarOperation {
	if op == nil {
		op = &xxx_ModifyCalendarOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.CalendarXML = o.CalendarXML
	op.Overwrite = o.Overwrite
	op.ChangeActivePolicy = o.ChangeActivePolicy
	return op
}

func (o *ModifyCalendarRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifyCalendarOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CalendarXML = op.CalendarXML
	o.Overwrite = op.Overwrite
	o.ChangeActivePolicy = op.ChangeActivePolicy
}
func (o *ModifyCalendarRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ModifyCalendarRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyCalendarOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyCalendarResponse structure represents the ModifyCalendar operation response
type ModifyCalendarResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ModifyCalendar return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ModifyCalendarResponse) xxx_ToOp(ctx context.Context, op *xxx_ModifyCalendarOperation) *xxx_ModifyCalendarOperation {
	if op == nil {
		op = &xxx_ModifyCalendarOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ModifyCalendarResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifyCalendarOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ModifyCalendarResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ModifyCalendarResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyCalendarOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteCalendarOperation structure represents the DeleteCalendar operation
type xxx_DeleteCalendarOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	CalendarName       *oaut.String   `idl:"name:bstrCalendarName" json:"calendar_name"`
	ChangeActivePolicy bool           `idl:"name:bChangeActivePolicy" json:"change_active_policy"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteCalendarOperation) OpNum() int { return 10 }

func (o *xxx_DeleteCalendarOperation) OpName() string { return "/IWRMCalendar/v0/DeleteCalendar" }

func (o *xxx_DeleteCalendarOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteCalendarOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CalendarName != nil {
			_ptr_bstrCalendarName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CalendarName != nil {
					if err := o.CalendarName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CalendarName, _ptr_bstrCalendarName); err != nil {
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
	// bChangeActivePolicy {in} (1:{alias=BOOL}(int32))
	{
		if !o.ChangeActivePolicy {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeleteCalendarOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrCalendarName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CalendarName == nil {
				o.CalendarName = &oaut.String{}
			}
			if err := o.CalendarName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrCalendarName := func(ptr interface{}) { o.CalendarName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CalendarName, _s_bstrCalendarName, _ptr_bstrCalendarName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bChangeActivePolicy {in} (1:{alias=BOOL}(int32))
	{
		var _bChangeActivePolicy int32
		if err := w.ReadData(&_bChangeActivePolicy); err != nil {
			return err
		}
		o.ChangeActivePolicy = _bChangeActivePolicy != 0
	}
	return nil
}

func (o *xxx_DeleteCalendarOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteCalendarOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteCalendarOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteCalendarRequest structure represents the DeleteCalendar operation request
type DeleteCalendarRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrCalendarName: A string that specifies the name of the calendar event.
	//
	// If this parameter is NULL, E_INVALIDARG MUST be returned.
	CalendarName *oaut.String `idl:"name:bstrCalendarName" json:"calendar_name"`
	// bChangeActivePolicy: A Boolean value that specifies whether the configuration changes
	// made by this method call SHOULD change the current active policy, if applicable,
	// of the system.<62>
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | The current active policy of the system SHOULD NOT be changed by the method      |
	//	|                  | call.                                                                            |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | If applicable, the current active policy of the system SHOULD be changed by the  |
	//	|                  | method call.                                                                     |
	//	+------------------+----------------------------------------------------------------------------------+
	ChangeActivePolicy bool `idl:"name:bChangeActivePolicy" json:"change_active_policy"`
}

func (o *DeleteCalendarRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteCalendarOperation) *xxx_DeleteCalendarOperation {
	if op == nil {
		op = &xxx_DeleteCalendarOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.CalendarName = o.CalendarName
	op.ChangeActivePolicy = o.ChangeActivePolicy
	return op
}

func (o *DeleteCalendarRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteCalendarOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CalendarName = op.CalendarName
	o.ChangeActivePolicy = op.ChangeActivePolicy
}
func (o *DeleteCalendarRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteCalendarRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteCalendarOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteCalendarResponse structure represents the DeleteCalendar operation response
type DeleteCalendarResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteCalendar return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteCalendarResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteCalendarOperation) *xxx_DeleteCalendarOperation {
	if op == nil {
		op = &xxx_DeleteCalendarOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteCalendarResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteCalendarOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteCalendarResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteCalendarResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteCalendarOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RenameCalendarOperation structure represents the RenameCalendar operation
type xxx_RenameCalendarOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	OldCalendarName *oaut.String   `idl:"name:bstrOldCalendarName" json:"old_calendar_name"`
	NewCalendarName *oaut.String   `idl:"name:bstrNewCalendarName" json:"new_calendar_name"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RenameCalendarOperation) OpNum() int { return 11 }

func (o *xxx_RenameCalendarOperation) OpName() string { return "/IWRMCalendar/v0/RenameCalendar" }

func (o *xxx_RenameCalendarOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameCalendarOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrOldCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.OldCalendarName != nil {
			_ptr_bstrOldCalendarName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OldCalendarName != nil {
					if err := o.OldCalendarName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OldCalendarName, _ptr_bstrOldCalendarName); err != nil {
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
	// bstrNewCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.NewCalendarName != nil {
			_ptr_bstrNewCalendarName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewCalendarName != nil {
					if err := o.NewCalendarName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewCalendarName, _ptr_bstrNewCalendarName); err != nil {
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

func (o *xxx_RenameCalendarOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrOldCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrOldCalendarName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OldCalendarName == nil {
				o.OldCalendarName = &oaut.String{}
			}
			if err := o.OldCalendarName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrOldCalendarName := func(ptr interface{}) { o.OldCalendarName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.OldCalendarName, _s_bstrOldCalendarName, _ptr_bstrOldCalendarName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrNewCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrNewCalendarName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewCalendarName == nil {
				o.NewCalendarName = &oaut.String{}
			}
			if err := o.NewCalendarName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrNewCalendarName := func(ptr interface{}) { o.NewCalendarName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.NewCalendarName, _s_bstrNewCalendarName, _ptr_bstrNewCalendarName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameCalendarOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameCalendarOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RenameCalendarOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RenameCalendarRequest structure represents the RenameCalendar operation request
type RenameCalendarRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrOldCalendarName: A string that specifies the current name of the calendar event
	// to be renamed.
	OldCalendarName *oaut.String `idl:"name:bstrOldCalendarName" json:"old_calendar_name"`
	// bstrNewCalendarName: A string that specifies the new name of the calendar event.
	NewCalendarName *oaut.String `idl:"name:bstrNewCalendarName" json:"new_calendar_name"`
}

func (o *RenameCalendarRequest) xxx_ToOp(ctx context.Context, op *xxx_RenameCalendarOperation) *xxx_RenameCalendarOperation {
	if op == nil {
		op = &xxx_RenameCalendarOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.OldCalendarName = o.OldCalendarName
	op.NewCalendarName = o.NewCalendarName
	return op
}

func (o *RenameCalendarRequest) xxx_FromOp(ctx context.Context, op *xxx_RenameCalendarOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.OldCalendarName = op.OldCalendarName
	o.NewCalendarName = op.NewCalendarName
}
func (o *RenameCalendarRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RenameCalendarRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameCalendarOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RenameCalendarResponse structure represents the RenameCalendar operation response
type RenameCalendarResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RenameCalendar return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RenameCalendarResponse) xxx_ToOp(ctx context.Context, op *xxx_RenameCalendarOperation) *xxx_RenameCalendarOperation {
	if op == nil {
		op = &xxx_RenameCalendarOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RenameCalendarResponse) xxx_FromOp(ctx context.Context, op *xxx_RenameCalendarOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RenameCalendarResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RenameCalendarResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameCalendarOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ComputeEventsOperation structure represents the ComputeEvents operation
type xxx_ComputeEventsOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	StartTime   *oaut.String   `idl:"name:szStartTime" json:"start_time"`
	EndTime     *oaut.String   `idl:"name:szEndTime" json:"end_time"`
	MergeEvents bool           `idl:"name:fMergeEvents" json:"merge_events"`
	Events      *oaut.String   `idl:"name:pbstrEvents" json:"events"`
	Conflicts   *oaut.String   `idl:"name:pbstrConflicts" json:"conflicts"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ComputeEventsOperation) OpNum() int { return 12 }

func (o *xxx_ComputeEventsOperation) OpName() string { return "/IWRMCalendar/v0/ComputeEvents" }

func (o *xxx_ComputeEventsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComputeEventsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szStartTime {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.StartTime != nil {
			_ptr_szStartTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.StartTime != nil {
					if err := o.StartTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.StartTime, _ptr_szStartTime); err != nil {
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
	// szEndTime {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EndTime != nil {
			_ptr_szEndTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EndTime != nil {
					if err := o.EndTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EndTime, _ptr_szEndTime); err != nil {
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
	// fMergeEvents {in} (1:{alias=BOOL}(int32))
	{
		if !o.MergeEvents {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ComputeEventsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szStartTime {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_szStartTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.StartTime == nil {
				o.StartTime = &oaut.String{}
			}
			if err := o.StartTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_szStartTime := func(ptr interface{}) { o.StartTime = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.StartTime, _s_szStartTime, _ptr_szStartTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// szEndTime {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_szEndTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EndTime == nil {
				o.EndTime = &oaut.String{}
			}
			if err := o.EndTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_szEndTime := func(ptr interface{}) { o.EndTime = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EndTime, _s_szEndTime, _ptr_szEndTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fMergeEvents {in} (1:{alias=BOOL}(int32))
	{
		var _bMergeEvents int32
		if err := w.ReadData(&_bMergeEvents); err != nil {
			return err
		}
		o.MergeEvents = _bMergeEvents != 0
	}
	return nil
}

func (o *xxx_ComputeEventsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComputeEventsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrEvents {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Events != nil {
			_ptr_pbstrEvents := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Events != nil {
					if err := o.Events.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Events, _ptr_pbstrEvents); err != nil {
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
	// pbstrConflicts {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Conflicts != nil {
			_ptr_pbstrConflicts := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Conflicts != nil {
					if err := o.Conflicts.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Conflicts, _ptr_pbstrConflicts); err != nil {
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

func (o *xxx_ComputeEventsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrEvents {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrEvents := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Events == nil {
				o.Events = &oaut.String{}
			}
			if err := o.Events.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrEvents := func(ptr interface{}) { o.Events = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Events, _s_pbstrEvents, _ptr_pbstrEvents); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrConflicts {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrConflicts := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Conflicts == nil {
				o.Conflicts = &oaut.String{}
			}
			if err := o.Conflicts.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrConflicts := func(ptr interface{}) { o.Conflicts = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Conflicts, _s_pbstrConflicts, _ptr_pbstrConflicts); err != nil {
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

// ComputeEventsRequest structure represents the ComputeEvents operation request
type ComputeEventsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szStartTime: A string that specifies the start time to compute calendar events, in
	// time format (section 2.2.1.3).
	StartTime *oaut.String `idl:"name:szStartTime" json:"start_time"`
	// szEndTime: A string that specifies the end time to compute calendar events, in time
	// format.
	EndTime *oaut.String `idl:"name:szEndTime" json:"end_time"`
	// fMergeEvents: A Boolean value that specifies whether to merge calendar events. All
	// conflicting events, that is, those whose start time and end time overlap, are returned
	// in pbstrConflicts. If the fMergeEvents value is TRUE, conflicting events are merged
	// into a new event with the start-time DtTmStart element picked from the event with
	// the earliest start time and the end-time DtTmEnd, PolicyName, and CalendarName elements
	// picked from the event with the latest end time.
	//
	//	+------------------+----------------------------------------------+
	//	|                  |                                              |
	//	|      VALUE       |                   MEANING                    |
	//	|                  |                                              |
	//	+------------------+----------------------------------------------+
	//	+------------------+----------------------------------------------+
	//	| FALSE 0x00000000 | The server SHOULD NOT merge calendar events. |
	//	+------------------+----------------------------------------------+
	//	| TRUE 0x00000001  | The server SHOULD merge calendar events.     |
	//	+------------------+----------------------------------------------+
	MergeEvents bool `idl:"name:fMergeEvents" json:"merge_events"`
}

func (o *ComputeEventsRequest) xxx_ToOp(ctx context.Context, op *xxx_ComputeEventsOperation) *xxx_ComputeEventsOperation {
	if op == nil {
		op = &xxx_ComputeEventsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.StartTime = o.StartTime
	op.EndTime = o.EndTime
	op.MergeEvents = o.MergeEvents
	return op
}

func (o *ComputeEventsRequest) xxx_FromOp(ctx context.Context, op *xxx_ComputeEventsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.StartTime = op.StartTime
	o.EndTime = op.EndTime
	o.MergeEvents = op.MergeEvents
}
func (o *ComputeEventsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ComputeEventsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ComputeEventsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ComputeEventsResponse structure represents the ComputeEvents operation response
type ComputeEventsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrEvents: A pointer to a string that returns a list of events that are computed
	// during execution, in the form of an Events element (section 2.2.5.15). For an example,
	// see Events Example (section 4.2.12).
	Events *oaut.String `idl:"name:pbstrEvents" json:"events"`
	// pbstrConflicts: A pointer to a string that returns a list of conflicting events during
	// the specified time interval, in the form of an Events element (section 2.2.5.15).
	// For an example, see Events Example (section 4.2.12).
	Conflicts *oaut.String `idl:"name:pbstrConflicts" json:"conflicts"`
	// Return: The ComputeEvents return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ComputeEventsResponse) xxx_ToOp(ctx context.Context, op *xxx_ComputeEventsOperation) *xxx_ComputeEventsOperation {
	if op == nil {
		op = &xxx_ComputeEventsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Events = o.Events
	op.Conflicts = o.Conflicts
	op.Return = o.Return
	return op
}

func (o *ComputeEventsResponse) xxx_FromOp(ctx context.Context, op *xxx_ComputeEventsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Events = op.Events
	o.Conflicts = op.Conflicts
	o.Return = op.Return
}
func (o *ComputeEventsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ComputeEventsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ComputeEventsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetScheduleInfoOperation structure represents the GetScheduleInfo operation
type xxx_GetScheduleInfoOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	ScheduleName *oaut.String   `idl:"name:bstrScheduleName" json:"schedule_name"`
	ScheduleXML  *oaut.String   `idl:"name:pbstrScheduleXML" json:"schedule_xml"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetScheduleInfoOperation) OpNum() int { return 13 }

func (o *xxx_GetScheduleInfoOperation) OpName() string { return "/IWRMCalendar/v0/GetScheduleInfo" }

func (o *xxx_GetScheduleInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetScheduleInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrScheduleName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ScheduleName != nil {
			_ptr_bstrScheduleName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ScheduleName != nil {
					if err := o.ScheduleName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ScheduleName, _ptr_bstrScheduleName); err != nil {
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

func (o *xxx_GetScheduleInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrScheduleName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrScheduleName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ScheduleName == nil {
				o.ScheduleName = &oaut.String{}
			}
			if err := o.ScheduleName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrScheduleName := func(ptr interface{}) { o.ScheduleName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ScheduleName, _s_bstrScheduleName, _ptr_bstrScheduleName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetScheduleInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetScheduleInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrScheduleXML {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ScheduleXML != nil {
			_ptr_pbstrScheduleXML := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ScheduleXML != nil {
					if err := o.ScheduleXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ScheduleXML, _ptr_pbstrScheduleXML); err != nil {
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

func (o *xxx_GetScheduleInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrScheduleXML {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrScheduleXML := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ScheduleXML == nil {
				o.ScheduleXML = &oaut.String{}
			}
			if err := o.ScheduleXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrScheduleXML := func(ptr interface{}) { o.ScheduleXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ScheduleXML, _s_pbstrScheduleXML, _ptr_pbstrScheduleXML); err != nil {
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

// GetScheduleInfoRequest structure represents the GetScheduleInfo operation request
type GetScheduleInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrScheduleName: A string that specifies the name of the schedule object for which
	// information is required.
	ScheduleName *oaut.String `idl:"name:bstrScheduleName" json:"schedule_name"`
}

func (o *GetScheduleInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetScheduleInfoOperation) *xxx_GetScheduleInfoOperation {
	if op == nil {
		op = &xxx_GetScheduleInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ScheduleName = o.ScheduleName
	return op
}

func (o *GetScheduleInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetScheduleInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ScheduleName = op.ScheduleName
}
func (o *GetScheduleInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetScheduleInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetScheduleInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetScheduleInfoResponse structure represents the GetScheduleInfo operation response
type GetScheduleInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrScheduleXML: A pointer to a string that returns the specified schedule object
	// information structure, in the form of a Schedule element (section 2.2.5.26). Sample
	// XML is provided in Schedule XML Example (section 4.2.22).
	ScheduleXML *oaut.String `idl:"name:pbstrScheduleXML" json:"schedule_xml"`
	// Return: The GetScheduleInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetScheduleInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetScheduleInfoOperation) *xxx_GetScheduleInfoOperation {
	if op == nil {
		op = &xxx_GetScheduleInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ScheduleXML = o.ScheduleXML
	op.Return = o.Return
	return op
}

func (o *GetScheduleInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetScheduleInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ScheduleXML = op.ScheduleXML
	o.Return = op.Return
}
func (o *GetScheduleInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetScheduleInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetScheduleInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateScheduleOperation structure represents the CreateSchedule operation
type xxx_CreateScheduleOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ScheduleXML *oaut.String   `idl:"name:bstrScheduleXML" json:"schedule_xml"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateScheduleOperation) OpNum() int { return 14 }

func (o *xxx_CreateScheduleOperation) OpName() string { return "/IWRMCalendar/v0/CreateSchedule" }

func (o *xxx_CreateScheduleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateScheduleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrScheduleXML {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ScheduleXML != nil {
			_ptr_bstrScheduleXML := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ScheduleXML != nil {
					if err := o.ScheduleXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ScheduleXML, _ptr_bstrScheduleXML); err != nil {
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

func (o *xxx_CreateScheduleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrScheduleXML {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrScheduleXML := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ScheduleXML == nil {
				o.ScheduleXML = &oaut.String{}
			}
			if err := o.ScheduleXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrScheduleXML := func(ptr interface{}) { o.ScheduleXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ScheduleXML, _s_bstrScheduleXML, _ptr_bstrScheduleXML); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateScheduleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateScheduleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateScheduleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateScheduleRequest structure represents the CreateSchedule operation request
type CreateScheduleRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrScheduleXML: A string that specifies the new schedule, in the form of a Schedule
	// element (section 2.2.5.26). Sample XML is provided in Schedule Example (section 4.2.22).
	ScheduleXML *oaut.String `idl:"name:bstrScheduleXML" json:"schedule_xml"`
}

func (o *CreateScheduleRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateScheduleOperation) *xxx_CreateScheduleOperation {
	if op == nil {
		op = &xxx_CreateScheduleOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ScheduleXML = o.ScheduleXML
	return op
}

func (o *CreateScheduleRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateScheduleOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ScheduleXML = op.ScheduleXML
}
func (o *CreateScheduleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateScheduleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateScheduleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateScheduleResponse structure represents the CreateSchedule operation response
type CreateScheduleResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateSchedule return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateScheduleResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateScheduleOperation) *xxx_CreateScheduleOperation {
	if op == nil {
		op = &xxx_CreateScheduleOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CreateScheduleResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateScheduleOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateScheduleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateScheduleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateScheduleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyScheduleOperation structure represents the ModifySchedule operation
type xxx_ModifyScheduleOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	ScheduleXML        *oaut.String   `idl:"name:bstrScheduleXML" json:"schedule_xml"`
	Overwrite          bool           `idl:"name:bOverwrite" json:"overwrite"`
	ChangeActivePolicy bool           `idl:"name:bChangeActivePolicy" json:"change_active_policy"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyScheduleOperation) OpNum() int { return 15 }

func (o *xxx_ModifyScheduleOperation) OpName() string { return "/IWRMCalendar/v0/ModifySchedule" }

func (o *xxx_ModifyScheduleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyScheduleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrScheduleXML {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ScheduleXML != nil {
			_ptr_bstrScheduleXML := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ScheduleXML != nil {
					if err := o.ScheduleXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ScheduleXML, _ptr_bstrScheduleXML); err != nil {
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
	// bOverwrite {in} (1:{alias=BOOL}(int32))
	{
		if !o.Overwrite {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// bChangeActivePolicy {in} (1:{alias=BOOL}(int32))
	{
		if !o.ChangeActivePolicy {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ModifyScheduleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrScheduleXML {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrScheduleXML := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ScheduleXML == nil {
				o.ScheduleXML = &oaut.String{}
			}
			if err := o.ScheduleXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrScheduleXML := func(ptr interface{}) { o.ScheduleXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ScheduleXML, _s_bstrScheduleXML, _ptr_bstrScheduleXML); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bOverwrite {in} (1:{alias=BOOL}(int32))
	{
		var _bOverwrite int32
		if err := w.ReadData(&_bOverwrite); err != nil {
			return err
		}
		o.Overwrite = _bOverwrite != 0
	}
	// bChangeActivePolicy {in} (1:{alias=BOOL}(int32))
	{
		var _bChangeActivePolicy int32
		if err := w.ReadData(&_bChangeActivePolicy); err != nil {
			return err
		}
		o.ChangeActivePolicy = _bChangeActivePolicy != 0
	}
	return nil
}

func (o *xxx_ModifyScheduleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyScheduleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ModifyScheduleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ModifyScheduleRequest structure represents the ModifySchedule operation request
type ModifyScheduleRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrScheduleXML: A string that specifies the modified schedule, in the form of a
	// Schedule element (section 2.2.5.26). Sample XML is provided in Schedule Example (section
	// 4.2.22).
	ScheduleXML *oaut.String `idl:"name:bstrScheduleXML" json:"schedule_xml"`
	// bOverwrite: A Boolean value that specifies whether to ignore the timestamp of the
	// specified schedule object when validating.
	//
	// A timestamp MUST be defined inside a common node at the root level of an XML element,
	// as shown in the Calendar example (section 4.2.6). The format of a timestamp is specified
	// in section 2.2.1.4.
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | The timestamp of the new schedule object MUST specify a time that is later       |
	//	|                  | than or equal to the timestamp of any modifications made to a calendar or        |
	//	|                  | schedule object on the server. Otherwise, the modification SHOULD fail, and      |
	//	|                  | WRM_ERR_OLD_INFORMATION SHOULD be returned.                                      |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | The schedule object is validated and modified without checking the timestamp.    |
	//	+------------------+----------------------------------------------------------------------------------+
	Overwrite bool `idl:"name:bOverwrite" json:"overwrite"`
	// bChangeActivePolicy: A Boolean value that specifies whether the configuration changes
	// made by this method call SHOULD change the current active policy, if applicable,
	// of the system.<66>
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | The current active policy of the system SHOULD NOT be changed by the method      |
	//	|                  | call.                                                                            |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | If applicable, the current active policy of the system SHOULD be changed by the  |
	//	|                  | method call.                                                                     |
	//	+------------------+----------------------------------------------------------------------------------+
	ChangeActivePolicy bool `idl:"name:bChangeActivePolicy" json:"change_active_policy"`
}

func (o *ModifyScheduleRequest) xxx_ToOp(ctx context.Context, op *xxx_ModifyScheduleOperation) *xxx_ModifyScheduleOperation {
	if op == nil {
		op = &xxx_ModifyScheduleOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ScheduleXML = o.ScheduleXML
	op.Overwrite = o.Overwrite
	op.ChangeActivePolicy = o.ChangeActivePolicy
	return op
}

func (o *ModifyScheduleRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifyScheduleOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ScheduleXML = op.ScheduleXML
	o.Overwrite = op.Overwrite
	o.ChangeActivePolicy = op.ChangeActivePolicy
}
func (o *ModifyScheduleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ModifyScheduleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyScheduleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyScheduleResponse structure represents the ModifySchedule operation response
type ModifyScheduleResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ModifySchedule return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ModifyScheduleResponse) xxx_ToOp(ctx context.Context, op *xxx_ModifyScheduleOperation) *xxx_ModifyScheduleOperation {
	if op == nil {
		op = &xxx_ModifyScheduleOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ModifyScheduleResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifyScheduleOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ModifyScheduleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ModifyScheduleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyScheduleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteScheduleOperation structure represents the DeleteSchedule operation
type xxx_DeleteScheduleOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	ScheduleName *oaut.String   `idl:"name:bstrScheduleName" json:"schedule_name"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteScheduleOperation) OpNum() int { return 16 }

func (o *xxx_DeleteScheduleOperation) OpName() string { return "/IWRMCalendar/v0/DeleteSchedule" }

func (o *xxx_DeleteScheduleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteScheduleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrScheduleName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ScheduleName != nil {
			_ptr_bstrScheduleName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ScheduleName != nil {
					if err := o.ScheduleName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ScheduleName, _ptr_bstrScheduleName); err != nil {
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

func (o *xxx_DeleteScheduleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrScheduleName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrScheduleName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ScheduleName == nil {
				o.ScheduleName = &oaut.String{}
			}
			if err := o.ScheduleName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrScheduleName := func(ptr interface{}) { o.ScheduleName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ScheduleName, _s_bstrScheduleName, _ptr_bstrScheduleName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteScheduleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteScheduleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteScheduleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteScheduleRequest structure represents the DeleteSchedule operation request
type DeleteScheduleRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrScheduleName: A string that specifies the name of the schedule object to be deleted.
	ScheduleName *oaut.String `idl:"name:bstrScheduleName" json:"schedule_name"`
}

func (o *DeleteScheduleRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteScheduleOperation) *xxx_DeleteScheduleOperation {
	if op == nil {
		op = &xxx_DeleteScheduleOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ScheduleName = o.ScheduleName
	return op
}

func (o *DeleteScheduleRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteScheduleOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ScheduleName = op.ScheduleName
}
func (o *DeleteScheduleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteScheduleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteScheduleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteScheduleResponse structure represents the DeleteSchedule operation response
type DeleteScheduleResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteSchedule return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteScheduleResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteScheduleOperation) *xxx_DeleteScheduleOperation {
	if op == nil {
		op = &xxx_DeleteScheduleOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteScheduleResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteScheduleOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteScheduleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteScheduleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteScheduleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RenameScheduleOperation structure represents the RenameSchedule operation
type xxx_RenameScheduleOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	OldScheduleName *oaut.String   `idl:"name:bstrOldScheduleName" json:"old_schedule_name"`
	NewScheduleName *oaut.String   `idl:"name:bstrNewScheduleName" json:"new_schedule_name"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RenameScheduleOperation) OpNum() int { return 17 }

func (o *xxx_RenameScheduleOperation) OpName() string { return "/IWRMCalendar/v0/RenameSchedule" }

func (o *xxx_RenameScheduleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameScheduleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrOldScheduleName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.OldScheduleName != nil {
			_ptr_bstrOldScheduleName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OldScheduleName != nil {
					if err := o.OldScheduleName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OldScheduleName, _ptr_bstrOldScheduleName); err != nil {
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
	// bstrNewScheduleName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.NewScheduleName != nil {
			_ptr_bstrNewScheduleName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewScheduleName != nil {
					if err := o.NewScheduleName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewScheduleName, _ptr_bstrNewScheduleName); err != nil {
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

func (o *xxx_RenameScheduleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrOldScheduleName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrOldScheduleName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OldScheduleName == nil {
				o.OldScheduleName = &oaut.String{}
			}
			if err := o.OldScheduleName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrOldScheduleName := func(ptr interface{}) { o.OldScheduleName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.OldScheduleName, _s_bstrOldScheduleName, _ptr_bstrOldScheduleName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrNewScheduleName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrNewScheduleName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewScheduleName == nil {
				o.NewScheduleName = &oaut.String{}
			}
			if err := o.NewScheduleName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrNewScheduleName := func(ptr interface{}) { o.NewScheduleName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.NewScheduleName, _s_bstrNewScheduleName, _ptr_bstrNewScheduleName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameScheduleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameScheduleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RenameScheduleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RenameScheduleRequest structure represents the RenameSchedule operation request
type RenameScheduleRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrOldScheduleName: A string that specifies the current name of the schedule object
	// to be renamed.
	OldScheduleName *oaut.String `idl:"name:bstrOldScheduleName" json:"old_schedule_name"`
	// bstrNewScheduleName: A string that specifies the new name of the schedule object.
	NewScheduleName *oaut.String `idl:"name:bstrNewScheduleName" json:"new_schedule_name"`
}

func (o *RenameScheduleRequest) xxx_ToOp(ctx context.Context, op *xxx_RenameScheduleOperation) *xxx_RenameScheduleOperation {
	if op == nil {
		op = &xxx_RenameScheduleOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.OldScheduleName = o.OldScheduleName
	op.NewScheduleName = o.NewScheduleName
	return op
}

func (o *RenameScheduleRequest) xxx_FromOp(ctx context.Context, op *xxx_RenameScheduleOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.OldScheduleName = op.OldScheduleName
	o.NewScheduleName = op.NewScheduleName
}
func (o *RenameScheduleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RenameScheduleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameScheduleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RenameScheduleResponse structure represents the RenameSchedule operation response
type RenameScheduleResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RenameSchedule return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RenameScheduleResponse) xxx_ToOp(ctx context.Context, op *xxx_RenameScheduleOperation) *xxx_RenameScheduleOperation {
	if op == nil {
		op = &xxx_RenameScheduleOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RenameScheduleResponse) xxx_FromOp(ctx context.Context, op *xxx_RenameScheduleOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RenameScheduleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RenameScheduleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameScheduleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MoveBeforeCalendarOperation structure represents the MoveBeforeCalendar operation
type xxx_MoveBeforeCalendarOperation struct {
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	CalendarName          *oaut.String   `idl:"name:bstrCalendarName" json:"calendar_name"`
	ReferenceCalendarName *oaut.String   `idl:"name:bstrRefCalendarName" json:"reference_calendar_name"`
	ChangeActivePolicy    bool           `idl:"name:bChangeActivePolicy" json:"change_active_policy"`
	Return                int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_MoveBeforeCalendarOperation) OpNum() int { return 18 }

func (o *xxx_MoveBeforeCalendarOperation) OpName() string {
	return "/IWRMCalendar/v0/MoveBeforeCalendar"
}

func (o *xxx_MoveBeforeCalendarOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveBeforeCalendarOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CalendarName != nil {
			_ptr_bstrCalendarName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CalendarName != nil {
					if err := o.CalendarName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CalendarName, _ptr_bstrCalendarName); err != nil {
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
	// bstrRefCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReferenceCalendarName != nil {
			_ptr_bstrRefCalendarName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferenceCalendarName != nil {
					if err := o.ReferenceCalendarName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferenceCalendarName, _ptr_bstrRefCalendarName); err != nil {
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
	// bChangeActivePolicy {in} (1:{alias=BOOL}(int32))
	{
		if !o.ChangeActivePolicy {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_MoveBeforeCalendarOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrCalendarName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CalendarName == nil {
				o.CalendarName = &oaut.String{}
			}
			if err := o.CalendarName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrCalendarName := func(ptr interface{}) { o.CalendarName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CalendarName, _s_bstrCalendarName, _ptr_bstrCalendarName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrRefCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrRefCalendarName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReferenceCalendarName == nil {
				o.ReferenceCalendarName = &oaut.String{}
			}
			if err := o.ReferenceCalendarName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrRefCalendarName := func(ptr interface{}) { o.ReferenceCalendarName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReferenceCalendarName, _s_bstrRefCalendarName, _ptr_bstrRefCalendarName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bChangeActivePolicy {in} (1:{alias=BOOL}(int32))
	{
		var _bChangeActivePolicy int32
		if err := w.ReadData(&_bChangeActivePolicy); err != nil {
			return err
		}
		o.ChangeActivePolicy = _bChangeActivePolicy != 0
	}
	return nil
}

func (o *xxx_MoveBeforeCalendarOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveBeforeCalendarOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MoveBeforeCalendarOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// MoveBeforeCalendarRequest structure represents the MoveBeforeCalendar operation request
type MoveBeforeCalendarRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrCalendarName: A string that specifies the name of the calendar event to be moved.
	CalendarName *oaut.String `idl:"name:bstrCalendarName" json:"calendar_name"`
	// bstrRefCalendarName: A string that specifies the name of a reference calendar event,
	// before which the specified event is to be moved.
	ReferenceCalendarName *oaut.String `idl:"name:bstrRefCalendarName" json:"reference_calendar_name"`
	// bChangeActivePolicy: A Boolean value that specifies whether the configuration changes
	// made by this method call SHOULD change the current active policy, if applicable,
	// of the system.<69>
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | The current active policy of the system SHOULD NOT be changed by the method      |
	//	|                  | call.                                                                            |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | If applicable, the current active policy of the system SHOULD be changed by the  |
	//	|                  | method call.                                                                     |
	//	+------------------+----------------------------------------------------------------------------------+
	ChangeActivePolicy bool `idl:"name:bChangeActivePolicy" json:"change_active_policy"`
}

func (o *MoveBeforeCalendarRequest) xxx_ToOp(ctx context.Context, op *xxx_MoveBeforeCalendarOperation) *xxx_MoveBeforeCalendarOperation {
	if op == nil {
		op = &xxx_MoveBeforeCalendarOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.CalendarName = o.CalendarName
	op.ReferenceCalendarName = o.ReferenceCalendarName
	op.ChangeActivePolicy = o.ChangeActivePolicy
	return op
}

func (o *MoveBeforeCalendarRequest) xxx_FromOp(ctx context.Context, op *xxx_MoveBeforeCalendarOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CalendarName = op.CalendarName
	o.ReferenceCalendarName = op.ReferenceCalendarName
	o.ChangeActivePolicy = op.ChangeActivePolicy
}
func (o *MoveBeforeCalendarRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MoveBeforeCalendarRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveBeforeCalendarOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MoveBeforeCalendarResponse structure represents the MoveBeforeCalendar operation response
type MoveBeforeCalendarResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MoveBeforeCalendar return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MoveBeforeCalendarResponse) xxx_ToOp(ctx context.Context, op *xxx_MoveBeforeCalendarOperation) *xxx_MoveBeforeCalendarOperation {
	if op == nil {
		op = &xxx_MoveBeforeCalendarOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *MoveBeforeCalendarResponse) xxx_FromOp(ctx context.Context, op *xxx_MoveBeforeCalendarOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *MoveBeforeCalendarResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MoveBeforeCalendarResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveBeforeCalendarOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MoveAfterCalendarOperation structure represents the MoveAfterCalendar operation
type xxx_MoveAfterCalendarOperation struct {
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	CalendarName          *oaut.String   `idl:"name:bstrCalendarName" json:"calendar_name"`
	ReferenceCalendarName *oaut.String   `idl:"name:bstrRefCalendarName" json:"reference_calendar_name"`
	ChangeActivePolicy    bool           `idl:"name:bChangeActivePolicy" json:"change_active_policy"`
	Return                int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_MoveAfterCalendarOperation) OpNum() int { return 19 }

func (o *xxx_MoveAfterCalendarOperation) OpName() string { return "/IWRMCalendar/v0/MoveAfterCalendar" }

func (o *xxx_MoveAfterCalendarOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveAfterCalendarOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CalendarName != nil {
			_ptr_bstrCalendarName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CalendarName != nil {
					if err := o.CalendarName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CalendarName, _ptr_bstrCalendarName); err != nil {
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
	// bstrRefCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReferenceCalendarName != nil {
			_ptr_bstrRefCalendarName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferenceCalendarName != nil {
					if err := o.ReferenceCalendarName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferenceCalendarName, _ptr_bstrRefCalendarName); err != nil {
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
	// bChangeActivePolicy {in} (1:{alias=BOOL}(int32))
	{
		if !o.ChangeActivePolicy {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_MoveAfterCalendarOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrCalendarName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CalendarName == nil {
				o.CalendarName = &oaut.String{}
			}
			if err := o.CalendarName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrCalendarName := func(ptr interface{}) { o.CalendarName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CalendarName, _s_bstrCalendarName, _ptr_bstrCalendarName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrRefCalendarName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrRefCalendarName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReferenceCalendarName == nil {
				o.ReferenceCalendarName = &oaut.String{}
			}
			if err := o.ReferenceCalendarName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrRefCalendarName := func(ptr interface{}) { o.ReferenceCalendarName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReferenceCalendarName, _s_bstrRefCalendarName, _ptr_bstrRefCalendarName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bChangeActivePolicy {in} (1:{alias=BOOL}(int32))
	{
		var _bChangeActivePolicy int32
		if err := w.ReadData(&_bChangeActivePolicy); err != nil {
			return err
		}
		o.ChangeActivePolicy = _bChangeActivePolicy != 0
	}
	return nil
}

func (o *xxx_MoveAfterCalendarOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveAfterCalendarOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MoveAfterCalendarOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// MoveAfterCalendarRequest structure represents the MoveAfterCalendar operation request
type MoveAfterCalendarRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrCalendarName: A string that specifies the name of the calendar event to be moved.
	CalendarName *oaut.String `idl:"name:bstrCalendarName" json:"calendar_name"`
	// bstrRefCalendarName: A string that specifies the name of a reference calendar event,
	// after which the specified event is to be moved.
	ReferenceCalendarName *oaut.String `idl:"name:bstrRefCalendarName" json:"reference_calendar_name"`
	// bChangeActivePolicy: A Boolean value that specifies whether the configuration changes
	// made by this method call SHOULD change the current active policy, if applicable,
	// of the system.<70>
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | The current active policy of the system SHOULD NOT be changed by the method      |
	//	|                  | call.                                                                            |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | If applicable, the current active policy of the system SHOULD be changed by the  |
	//	|                  | method call.                                                                     |
	//	+------------------+----------------------------------------------------------------------------------+
	ChangeActivePolicy bool `idl:"name:bChangeActivePolicy" json:"change_active_policy"`
}

func (o *MoveAfterCalendarRequest) xxx_ToOp(ctx context.Context, op *xxx_MoveAfterCalendarOperation) *xxx_MoveAfterCalendarOperation {
	if op == nil {
		op = &xxx_MoveAfterCalendarOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.CalendarName = o.CalendarName
	op.ReferenceCalendarName = o.ReferenceCalendarName
	op.ChangeActivePolicy = o.ChangeActivePolicy
	return op
}

func (o *MoveAfterCalendarRequest) xxx_FromOp(ctx context.Context, op *xxx_MoveAfterCalendarOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CalendarName = op.CalendarName
	o.ReferenceCalendarName = op.ReferenceCalendarName
	o.ChangeActivePolicy = op.ChangeActivePolicy
}
func (o *MoveAfterCalendarRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MoveAfterCalendarRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveAfterCalendarOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MoveAfterCalendarResponse structure represents the MoveAfterCalendar operation response
type MoveAfterCalendarResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MoveAfterCalendar return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MoveAfterCalendarResponse) xxx_ToOp(ctx context.Context, op *xxx_MoveAfterCalendarOperation) *xxx_MoveAfterCalendarOperation {
	if op == nil {
		op = &xxx_MoveAfterCalendarOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *MoveAfterCalendarResponse) xxx_FromOp(ctx context.Context, op *xxx_MoveAfterCalendarOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *MoveAfterCalendarResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MoveAfterCalendarResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveAfterCalendarOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetServerTimeZoneOperation structure represents the GetServerTimeZone operation
type xxx_GetServerTimeZoneOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	ServerTimeZone int32          `idl:"name:pnServerTimeZone" json:"server_time_zone"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetServerTimeZoneOperation) OpNum() int { return 20 }

func (o *xxx_GetServerTimeZoneOperation) OpName() string { return "/IWRMCalendar/v0/GetServerTimeZone" }

func (o *xxx_GetServerTimeZoneOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerTimeZoneOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetServerTimeZoneOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetServerTimeZoneOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerTimeZoneOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pnServerTimeZone {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.ServerTimeZone); err != nil {
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

func (o *xxx_GetServerTimeZoneOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pnServerTimeZone {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.ServerTimeZone); err != nil {
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

// GetServerTimeZoneRequest structure represents the GetServerTimeZone operation request
type GetServerTimeZoneRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetServerTimeZoneRequest) xxx_ToOp(ctx context.Context, op *xxx_GetServerTimeZoneOperation) *xxx_GetServerTimeZoneOperation {
	if op == nil {
		op = &xxx_GetServerTimeZoneOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetServerTimeZoneRequest) xxx_FromOp(ctx context.Context, op *xxx_GetServerTimeZoneOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetServerTimeZoneRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetServerTimeZoneRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerTimeZoneOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetServerTimeZoneResponse structure represents the GetServerTimeZone operation response
type GetServerTimeZoneResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pnServerTimeZone: A pointer to a 32-bit signed integer that returns the server time
	// zone.
	ServerTimeZone int32 `idl:"name:pnServerTimeZone" json:"server_time_zone"`
	// Return: The GetServerTimeZone return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetServerTimeZoneResponse) xxx_ToOp(ctx context.Context, op *xxx_GetServerTimeZoneOperation) *xxx_GetServerTimeZoneOperation {
	if op == nil {
		op = &xxx_GetServerTimeZoneOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ServerTimeZone = o.ServerTimeZone
	op.Return = o.Return
	return op
}

func (o *GetServerTimeZoneResponse) xxx_FromOp(ctx context.Context, op *xxx_GetServerTimeZoneOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ServerTimeZone = op.ServerTimeZone
	o.Return = op.Return
}
func (o *GetServerTimeZoneResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetServerTimeZoneResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerTimeZoneOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
