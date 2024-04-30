package ialertdatacollector

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idatacollector "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/idatacollector/v0"
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
	_ = idatacollector.GoPackage
)

// IAlertDataCollector server interface.
type AlertDataCollectorServer interface {

	// IDataCollector base class.
	idatacollector.DataCollectorServer

	// AlertThresholds operation.
	GetAlertThresholds(context.Context, *GetAlertThresholdsRequest) (*GetAlertThresholdsResponse, error)

	// AlertThresholds operation.
	SetAlertThresholds(context.Context, *SetAlertThresholdsRequest) (*SetAlertThresholdsResponse, error)

	// EventLog operation.
	GetEventLog(context.Context, *GetEventLogRequest) (*GetEventLogResponse, error)

	// EventLog operation.
	SetEventLog(context.Context, *SetEventLogRequest) (*SetEventLogResponse, error)

	// SampleInterval operation.
	GetSampleInterval(context.Context, *GetSampleIntervalRequest) (*GetSampleIntervalResponse, error)

	// SampleInterval operation.
	SetSampleInterval(context.Context, *SetSampleIntervalRequest) (*SetSampleIntervalResponse, error)

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

	// TriggerDataCollectorSet operation.
	GetTriggerDataCollectorSet(context.Context, *GetTriggerDataCollectorSetRequest) (*GetTriggerDataCollectorSetResponse, error)

	// TriggerDataCollectorSet operation.
	SetTriggerDataCollectorSet(context.Context, *SetTriggerDataCollectorSetRequest) (*SetTriggerDataCollectorSetResponse, error)
}

func RegisterAlertDataCollectorServer(conn dcerpc.Conn, o AlertDataCollectorServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAlertDataCollectorServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AlertDataCollectorSyntaxV0_0))...)
}

func NewAlertDataCollectorServerHandle(o AlertDataCollectorServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AlertDataCollectorServerHandle(ctx, o, opNum, r)
	}
}

func AlertDataCollectorServerHandle(ctx context.Context, o AlertDataCollectorServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 32 {
		// IDataCollector base method.
		return idatacollector.DataCollectorServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 32: // AlertThresholds
		in := &GetAlertThresholdsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAlertThresholds(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // AlertThresholds
		in := &SetAlertThresholdsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetAlertThresholds(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // EventLog
		in := &GetEventLogRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventLog(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // EventLog
		in := &SetEventLogRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEventLog(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // SampleInterval
		in := &GetSampleIntervalRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSampleInterval(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // SampleInterval
		in := &SetSampleIntervalRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSampleInterval(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // Task
		in := &GetTaskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTask(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // Task
		in := &SetTaskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetTask(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // TaskRunAsSelf
		in := &GetTaskRunAsSelfRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTaskRunAsSelf(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // TaskRunAsSelf
		in := &SetTaskRunAsSelfRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetTaskRunAsSelf(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // TaskArguments
		in := &GetTaskArgumentsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTaskArguments(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // TaskArguments
		in := &SetTaskArgumentsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetTaskArguments(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // TaskUserTextArguments
		in := &GetTaskUserTextArgumentsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTaskUserTextArguments(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // TaskUserTextArguments
		in := &SetTaskUserTextArgumentsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetTaskUserTextArguments(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // TriggerDataCollectorSet
		in := &GetTriggerDataCollectorSetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTriggerDataCollectorSet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // TriggerDataCollectorSet
		in := &SetTriggerDataCollectorSetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetTriggerDataCollectorSet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
