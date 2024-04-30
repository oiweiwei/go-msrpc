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
		in := &GetStartDateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetStartDate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // StartDate
		in := &SetStartDateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetStartDate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // EndDate
		in := &GetEndDateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEndDate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // EndDate
		in := &SetEndDateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEndDate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // StartTime
		in := &GetStartTimeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetStartTime(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // StartTime
		in := &SetStartTimeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetStartTime(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // Days
		in := &GetDaysRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDays(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // Days
		in := &SetDaysRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDays(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
