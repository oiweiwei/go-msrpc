package ischedule

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

// ISchedule server interface.
type ScheduleServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// StartDate operation.
	GetStartDate(context.Context, *GetStartDateRequest) (*GetStartDateResponse, error)

	// StartDate operation.
	SetStartDate(context.Context, *SetStartDateRequest) (*SetStartDateResponse, error)

	// EndDate operation.
	GetEndDate(context.Context, *GetEndDateRequest) (*GetEndDateResponse, error)

	// EndDate operation.
	SetEndDate(context.Context, *SetEndDateRequest) (*SetEndDateResponse, error)

	// StartTime operation.
	GetStartTime(context.Context, *GetStartTimeRequest) (*GetStartTimeResponse, error)

	// StartTime operation.
	SetStartTime(context.Context, *SetStartTimeRequest) (*SetStartTimeResponse, error)

	// Days operation.
	GetDays(context.Context, *GetDaysRequest) (*GetDaysResponse, error)

	// Days operation.
	SetDays(context.Context, *SetDaysRequest) (*SetDaysResponse, error)
}

func RegisterScheduleServer(conn dcerpc.Conn, o ScheduleServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewScheduleServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ScheduleSyntaxV0_0))...)
}

func NewScheduleServerHandle(o ScheduleServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ScheduleServerHandle(ctx, o, opNum, r)
	}
}

func ScheduleServerHandle(ctx context.Context, o ScheduleServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // StartDate
		op := &xxx_GetStartDateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetStartDateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetStartDate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // StartDate
		op := &xxx_SetStartDateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetStartDateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetStartDate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // EndDate
		op := &xxx_GetEndDateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEndDateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEndDate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // EndDate
		op := &xxx_SetEndDateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEndDateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEndDate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // StartTime
		op := &xxx_GetStartTimeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetStartTimeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetStartTime(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // StartTime
		op := &xxx_SetStartTimeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetStartTimeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetStartTime(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Days
		op := &xxx_GetDaysOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDaysRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDays(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Days
		op := &xxx_SetDaysOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDaysRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDays(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ISchedule
type UnimplementedScheduleServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedScheduleServer) GetStartDate(context.Context, *GetStartDateRequest) (*GetStartDateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedScheduleServer) SetStartDate(context.Context, *SetStartDateRequest) (*SetStartDateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedScheduleServer) GetEndDate(context.Context, *GetEndDateRequest) (*GetEndDateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedScheduleServer) SetEndDate(context.Context, *SetEndDateRequest) (*SetEndDateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedScheduleServer) GetStartTime(context.Context, *GetStartTimeRequest) (*GetStartTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedScheduleServer) SetStartTime(context.Context, *SetStartTimeRequest) (*SetStartTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedScheduleServer) GetDays(context.Context, *GetDaysRequest) (*GetDaysResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedScheduleServer) SetDays(context.Context, *SetDaysRequest) (*SetDaysResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ScheduleServer = (*UnimplementedScheduleServer)(nil)
