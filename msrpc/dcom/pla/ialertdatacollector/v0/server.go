package ialertdatacollector

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
		op := &xxx_GetAlertThresholdsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAlertThresholdsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAlertThresholds(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // AlertThresholds
		op := &xxx_SetAlertThresholdsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAlertThresholdsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAlertThresholds(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // EventLog
		op := &xxx_GetEventLogOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEventLogRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEventLog(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // EventLog
		op := &xxx_SetEventLogOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEventLogRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEventLog(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // SampleInterval
		op := &xxx_GetSampleIntervalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSampleIntervalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSampleInterval(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // SampleInterval
		op := &xxx_SetSampleIntervalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSampleIntervalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSampleInterval(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // Task
		op := &xxx_GetTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // Task
		op := &xxx_SetTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // TaskRunAsSelf
		op := &xxx_GetTaskRunAsSelfOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTaskRunAsSelfRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTaskRunAsSelf(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // TaskRunAsSelf
		op := &xxx_SetTaskRunAsSelfOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTaskRunAsSelfRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTaskRunAsSelf(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // TaskArguments
		op := &xxx_GetTaskArgumentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTaskArgumentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTaskArguments(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // TaskArguments
		op := &xxx_SetTaskArgumentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTaskArgumentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTaskArguments(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // TaskUserTextArguments
		op := &xxx_GetTaskUserTextArgumentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTaskUserTextArgumentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTaskUserTextArguments(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // TaskUserTextArguments
		op := &xxx_SetTaskUserTextArgumentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTaskUserTextArgumentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTaskUserTextArguments(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // TriggerDataCollectorSet
		op := &xxx_GetTriggerDataCollectorSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTriggerDataCollectorSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTriggerDataCollectorSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // TriggerDataCollectorSet
		op := &xxx_SetTriggerDataCollectorSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTriggerDataCollectorSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTriggerDataCollectorSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAlertDataCollector
type UnimplementedAlertDataCollectorServer struct {
	idatacollector.UnimplementedDataCollectorServer
}

func (UnimplementedAlertDataCollectorServer) GetAlertThresholds(context.Context, *GetAlertThresholdsRequest) (*GetAlertThresholdsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) SetAlertThresholds(context.Context, *SetAlertThresholdsRequest) (*SetAlertThresholdsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) GetEventLog(context.Context, *GetEventLogRequest) (*GetEventLogResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) SetEventLog(context.Context, *SetEventLogRequest) (*SetEventLogResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) GetSampleInterval(context.Context, *GetSampleIntervalRequest) (*GetSampleIntervalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) SetSampleInterval(context.Context, *SetSampleIntervalRequest) (*SetSampleIntervalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) SetTask(context.Context, *SetTaskRequest) (*SetTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) GetTaskRunAsSelf(context.Context, *GetTaskRunAsSelfRequest) (*GetTaskRunAsSelfResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) SetTaskRunAsSelf(context.Context, *SetTaskRunAsSelfRequest) (*SetTaskRunAsSelfResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) GetTaskArguments(context.Context, *GetTaskArgumentsRequest) (*GetTaskArgumentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) SetTaskArguments(context.Context, *SetTaskArgumentsRequest) (*SetTaskArgumentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) GetTaskUserTextArguments(context.Context, *GetTaskUserTextArgumentsRequest) (*GetTaskUserTextArgumentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) SetTaskUserTextArguments(context.Context, *SetTaskUserTextArgumentsRequest) (*SetTaskUserTextArgumentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) GetTriggerDataCollectorSet(context.Context, *GetTriggerDataCollectorSetRequest) (*GetTriggerDataCollectorSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAlertDataCollectorServer) SetTriggerDataCollectorSet(context.Context, *SetTriggerDataCollectorSetRequest) (*SetTriggerDataCollectorSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AlertDataCollectorServer = (*UnimplementedAlertDataCollectorServer)(nil)
