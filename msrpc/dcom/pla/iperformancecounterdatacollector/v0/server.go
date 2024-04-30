package iperformancecounterdatacollector

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
		in := &GetDataSourceNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDataSourceName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // DataSourceName
		in := &SetDataSourceNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDataSourceName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // PerformanceCounters
		in := &GetPerformanceCountersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPerformanceCounters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // PerformanceCounters
		in := &SetPerformanceCountersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPerformanceCounters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // LogFileFormat
		in := &GetLogFileFormatRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLogFileFormat(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // LogFileFormat
		in := &SetLogFileFormatRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetLogFileFormat(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // SampleInterval
		in := &GetSampleIntervalRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSampleInterval(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // SampleInterval
		in := &SetSampleIntervalRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSampleInterval(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // SegmentMaxRecords
		in := &GetSegmentMaxRecordsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSegmentMaxRecords(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // SegmentMaxRecords
		in := &SetSegmentMaxRecordsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSegmentMaxRecords(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
