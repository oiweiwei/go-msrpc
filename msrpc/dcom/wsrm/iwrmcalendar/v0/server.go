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

	GetCalendarInfo(context.Context, *GetCalendarInfoRequest) (*GetCalendarInfoResponse, error)

	CreateCalendar(context.Context, *CreateCalendarRequest) (*CreateCalendarResponse, error)

	ModifyCalendar(context.Context, *ModifyCalendarRequest) (*ModifyCalendarResponse, error)

	DeleteCalendar(context.Context, *DeleteCalendarRequest) (*DeleteCalendarResponse, error)

	RenameCalendar(context.Context, *RenameCalendarRequest) (*RenameCalendarResponse, error)

	ComputeEvents(context.Context, *ComputeEventsRequest) (*ComputeEventsResponse, error)

	GetScheduleInfo(context.Context, *GetScheduleInfoRequest) (*GetScheduleInfoResponse, error)

	CreateSchedule(context.Context, *CreateScheduleRequest) (*CreateScheduleResponse, error)

	ModifySchedule(context.Context, *ModifyScheduleRequest) (*ModifyScheduleResponse, error)

	DeleteSchedule(context.Context, *DeleteScheduleRequest) (*DeleteScheduleResponse, error)

	RenameSchedule(context.Context, *RenameScheduleRequest) (*RenameScheduleResponse, error)

	MoveBeforeCalendar(context.Context, *MoveBeforeCalendarRequest) (*MoveBeforeCalendarResponse, error)

	MoveAfterCalendar(context.Context, *MoveAfterCalendarRequest) (*MoveAfterCalendarResponse, error)

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
