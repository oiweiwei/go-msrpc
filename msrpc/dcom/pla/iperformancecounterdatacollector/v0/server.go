package iperformancecounterdatacollector

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

// IPerformanceCounterDataCollector server interface.
type PerformanceCounterDataCollectorServer interface {

	// IDataCollector base class.
	idatacollector.DataCollectorServer

	// DataSourceName operation.
	GetDataSourceName(context.Context, *GetDataSourceNameRequest) (*GetDataSourceNameResponse, error)

	// DataSourceName operation.
	SetDataSourceName(context.Context, *SetDataSourceNameRequest) (*SetDataSourceNameResponse, error)

	// PerformanceCounters operation.
	GetPerformanceCounters(context.Context, *GetPerformanceCountersRequest) (*GetPerformanceCountersResponse, error)

	// PerformanceCounters operation.
	SetPerformanceCounters(context.Context, *SetPerformanceCountersRequest) (*SetPerformanceCountersResponse, error)

	// LogFileFormat operation.
	GetLogFileFormat(context.Context, *GetLogFileFormatRequest) (*GetLogFileFormatResponse, error)

	// LogFileFormat operation.
	SetLogFileFormat(context.Context, *SetLogFileFormatRequest) (*SetLogFileFormatResponse, error)

	// SampleInterval operation.
	GetSampleInterval(context.Context, *GetSampleIntervalRequest) (*GetSampleIntervalResponse, error)

	// SampleInterval operation.
	SetSampleInterval(context.Context, *SetSampleIntervalRequest) (*SetSampleIntervalResponse, error)

	// SegmentMaxRecords operation.
	GetSegmentMaxRecords(context.Context, *GetSegmentMaxRecordsRequest) (*GetSegmentMaxRecordsResponse, error)

	// SegmentMaxRecords operation.
	SetSegmentMaxRecords(context.Context, *SetSegmentMaxRecordsRequest) (*SetSegmentMaxRecordsResponse, error)
}

func RegisterPerformanceCounterDataCollectorServer(conn dcerpc.Conn, o PerformanceCounterDataCollectorServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPerformanceCounterDataCollectorServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PerformanceCounterDataCollectorSyntaxV0_0))...)
}

func NewPerformanceCounterDataCollectorServerHandle(o PerformanceCounterDataCollectorServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PerformanceCounterDataCollectorServerHandle(ctx, o, opNum, r)
	}
}

func PerformanceCounterDataCollectorServerHandle(ctx context.Context, o PerformanceCounterDataCollectorServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 32 {
		// IDataCollector base method.
		return idatacollector.DataCollectorServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 32: // DataSourceName
		op := &xxx_GetDataSourceNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDataSourceNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDataSourceName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // DataSourceName
		op := &xxx_SetDataSourceNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDataSourceNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDataSourceName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // PerformanceCounters
		op := &xxx_GetPerformanceCountersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPerformanceCountersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPerformanceCounters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // PerformanceCounters
		op := &xxx_SetPerformanceCountersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPerformanceCountersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPerformanceCounters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // LogFileFormat
		op := &xxx_GetLogFileFormatOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLogFileFormatRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLogFileFormat(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // LogFileFormat
		op := &xxx_SetLogFileFormatOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLogFileFormatRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLogFileFormat(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // SampleInterval
		op := &xxx_GetSampleIntervalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSampleIntervalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSampleInterval(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // SampleInterval
		op := &xxx_SetSampleIntervalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSampleIntervalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSampleInterval(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // SegmentMaxRecords
		op := &xxx_GetSegmentMaxRecordsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSegmentMaxRecordsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSegmentMaxRecords(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // SegmentMaxRecords
		op := &xxx_SetSegmentMaxRecordsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSegmentMaxRecordsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSegmentMaxRecords(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IPerformanceCounterDataCollector
type UnimplementedPerformanceCounterDataCollectorServer struct {
	idatacollector.UnimplementedDataCollectorServer
}

func (UnimplementedPerformanceCounterDataCollectorServer) GetDataSourceName(context.Context, *GetDataSourceNameRequest) (*GetDataSourceNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerformanceCounterDataCollectorServer) SetDataSourceName(context.Context, *SetDataSourceNameRequest) (*SetDataSourceNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerformanceCounterDataCollectorServer) GetPerformanceCounters(context.Context, *GetPerformanceCountersRequest) (*GetPerformanceCountersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerformanceCounterDataCollectorServer) SetPerformanceCounters(context.Context, *SetPerformanceCountersRequest) (*SetPerformanceCountersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerformanceCounterDataCollectorServer) GetLogFileFormat(context.Context, *GetLogFileFormatRequest) (*GetLogFileFormatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerformanceCounterDataCollectorServer) SetLogFileFormat(context.Context, *SetLogFileFormatRequest) (*SetLogFileFormatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerformanceCounterDataCollectorServer) GetSampleInterval(context.Context, *GetSampleIntervalRequest) (*GetSampleIntervalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerformanceCounterDataCollectorServer) SetSampleInterval(context.Context, *SetSampleIntervalRequest) (*SetSampleIntervalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerformanceCounterDataCollectorServer) GetSegmentMaxRecords(context.Context, *GetSegmentMaxRecordsRequest) (*GetSegmentMaxRecordsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerformanceCounterDataCollectorServer) SetSegmentMaxRecords(context.Context, *SetSegmentMaxRecordsRequest) (*SetSegmentMaxRecordsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PerformanceCounterDataCollectorServer = (*UnimplementedPerformanceCounterDataCollectorServer)(nil)
