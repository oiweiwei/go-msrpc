package itracedatacollector

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
	pla "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla"
	idatacollector "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/idatacollector/v0"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = idatacollector.GoPackage
	_ = pla.GoPackage
	_ = dtyp.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/pla"
)

var (
	// ITraceDataCollector interface identifier 0383750b-098b-11d8-9414-505054503030
	TraceDataCollectorIID = &dcom.IID{Data1: 0x0383750b, Data2: 0x098b, Data3: 0x11d8, Data4: []byte{0x94, 0x14, 0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	// Syntax UUID
	TraceDataCollectorSyntaxUUID = &uuid.UUID{TimeLow: 0x383750b, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	// Syntax ID
	TraceDataCollectorSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: TraceDataCollectorSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ITraceDataCollector interface.
type TraceDataCollectorClient interface {

	// IDataCollector retrieval method.
	DataCollector() idatacollector.DataCollectorClient

	// BufferSize operation.
	GetBufferSize(context.Context, *GetBufferSizeRequest, ...dcerpc.CallOption) (*GetBufferSizeResponse, error)

	// BufferSize operation.
	SetBufferSize(context.Context, *SetBufferSizeRequest, ...dcerpc.CallOption) (*SetBufferSizeResponse, error)

	// The BuffersLost (Get) method retrieves the BufferLost property, as specified in the
	// property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetBuffersLost(context.Context, *GetBuffersLostRequest, ...dcerpc.CallOption) (*GetBuffersLostResponse, error)

	// Opnum35NotUsedOnWire operation.
	// Opnum35NotUsedOnWire

	// The BuffersWritten (Get) method retrieves the BuffersWritten property, as specified
	// in the property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetBuffersWritten(context.Context, *GetBuffersWrittenRequest, ...dcerpc.CallOption) (*GetBuffersWrittenResponse, error)

	// Opnum37NotUsedOnWire operation.
	// Opnum37NotUsedOnWire

	// The ClockType enumeration defines the clock resolution to use when tracing events.
	GetClockType(context.Context, *GetClockTypeRequest, ...dcerpc.CallOption) (*GetClockTypeResponse, error)

	// The ClockType enumeration defines the clock resolution to use when tracing events.
	SetClockType(context.Context, *SetClockTypeRequest, ...dcerpc.CallOption) (*SetClockTypeResponse, error)

	// The EventsLost (Get) method retrieves the EventsLost property, as specified in the
	// property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetEventsLost(context.Context, *GetEventsLostRequest, ...dcerpc.CallOption) (*GetEventsLostResponse, error)

	// Opnum41NotUsedOnWire operation.
	// Opnum41NotUsedOnWire

	// The ExtendedModes (Put) method sets the ExtendedModes property, as specified in the
	// property table in section 3.2.4.9.
	//
	// The ExtendedModes (Get) method retrieves the ExtendedModes property, as specified
	// in the property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetExtendedModes(context.Context, *GetExtendedModesRequest, ...dcerpc.CallOption) (*GetExtendedModesResponse, error)

	// The ExtendedModes (Put) method sets the ExtendedModes property, as specified in the
	// property table in section 3.2.4.9.
	//
	// The ExtendedModes (Get) method retrieves the ExtendedModes property, as specified
	// in the property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	SetExtendedModes(context.Context, *SetExtendedModesRequest, ...dcerpc.CallOption) (*SetExtendedModesResponse, error)

	// FlushTimer operation.
	GetFlushTimer(context.Context, *GetFlushTimerRequest, ...dcerpc.CallOption) (*GetFlushTimerResponse, error)

	// FlushTimer operation.
	SetFlushTimer(context.Context, *SetFlushTimerRequest, ...dcerpc.CallOption) (*SetFlushTimerResponse, error)

	// The FreeBuffers (Get) method retrieves the FreeBuffers property, as specified in
	// the property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetFreeBuffers(context.Context, *GetFreeBuffersRequest, ...dcerpc.CallOption) (*GetFreeBuffersResponse, error)

	// Opnum47NotUsedOnWire operation.
	// Opnum47NotUsedOnWire

	// Guid operation.
	GetGUID(context.Context, *GetGUIDRequest, ...dcerpc.CallOption) (*GetGUIDResponse, error)

	// Guid operation.
	SetGUID(context.Context, *SetGUIDRequest, ...dcerpc.CallOption) (*SetGUIDResponse, error)

	// The IsKernelTrace (Get) method retrieves the IsKernelTrace property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetIsKernelTrace(context.Context, *GetIsKernelTraceRequest, ...dcerpc.CallOption) (*GetIsKernelTraceResponse, error)

	// MaximumBuffers operation.
	GetMaximumBuffers(context.Context, *GetMaximumBuffersRequest, ...dcerpc.CallOption) (*GetMaximumBuffersResponse, error)

	// MaximumBuffers operation.
	SetMaximumBuffers(context.Context, *SetMaximumBuffersRequest, ...dcerpc.CallOption) (*SetMaximumBuffersResponse, error)

	// MinimumBuffers operation.
	GetMinimumBuffers(context.Context, *GetMinimumBuffersRequest, ...dcerpc.CallOption) (*GetMinimumBuffersResponse, error)

	// MinimumBuffers operation.
	SetMinimumBuffers(context.Context, *SetMinimumBuffersRequest, ...dcerpc.CallOption) (*SetMinimumBuffersResponse, error)

	// NumberOfBuffers operation.
	GetNumberOfBuffers(context.Context, *GetNumberOfBuffersRequest, ...dcerpc.CallOption) (*GetNumberOfBuffersResponse, error)

	// NumberOfBuffers operation.
	SetNumberOfBuffers(context.Context, *SetNumberOfBuffersRequest, ...dcerpc.CallOption) (*SetNumberOfBuffersResponse, error)

	// PreallocateFile operation.
	GetPreallocateFile(context.Context, *GetPreallocateFileRequest, ...dcerpc.CallOption) (*GetPreallocateFileResponse, error)

	// PreallocateFile operation.
	SetPreallocateFile(context.Context, *SetPreallocateFileRequest, ...dcerpc.CallOption) (*SetPreallocateFileResponse, error)

	// ProcessMode operation.
	GetProcessMode(context.Context, *GetProcessModeRequest, ...dcerpc.CallOption) (*GetProcessModeResponse, error)

	// ProcessMode operation.
	SetProcessMode(context.Context, *SetProcessModeRequest, ...dcerpc.CallOption) (*SetProcessModeResponse, error)

	// The RealTimeBuffersLost (Get) method retrieves the RealTimeBuffersLost property,
	// as specified in the property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetRealTimeBuffersLost(context.Context, *GetRealTimeBuffersLostRequest, ...dcerpc.CallOption) (*GetRealTimeBuffersLostResponse, error)

	// Opnum62NotUsedOnWire operation.
	// Opnum62NotUsedOnWire

	// The SessionId (Get) method retrieves the SessionId property, as specified in the
	// property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetSessionID(context.Context, *GetSessionIDRequest, ...dcerpc.CallOption) (*GetSessionIDResponse, error)

	// Opnum64NotUsedOnWire operation.
	// Opnum64NotUsedOnWire

	// SessionName operation.
	GetSessionName(context.Context, *GetSessionNameRequest, ...dcerpc.CallOption) (*GetSessionNameResponse, error)

	// SessionName operation.
	SetSessionName(context.Context, *SetSessionNameRequest, ...dcerpc.CallOption) (*SetSessionNameResponse, error)

	// The SessionThreadId (Get) method retrieves the SessionThreadId property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetSessionThreadID(context.Context, *GetSessionThreadIDRequest, ...dcerpc.CallOption) (*GetSessionThreadIDResponse, error)

	// Opnum68NotUsedOnWire operation.
	// Opnum68NotUsedOnWire

	// The StreamMode enumeration defines where the trace events are delivered.
	GetStreamMode(context.Context, *GetStreamModeRequest, ...dcerpc.CallOption) (*GetStreamModeResponse, error)

	// The StreamMode enumeration defines where the trace events are delivered.
	SetStreamMode(context.Context, *SetStreamModeRequest, ...dcerpc.CallOption) (*SetStreamModeResponse, error)

	// The TraceDataProviders (Get) method retrieves the TraceDataProviders property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetTraceDataProviders(context.Context, *GetTraceDataProvidersRequest, ...dcerpc.CallOption) (*GetTraceDataProvidersResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) TraceDataCollectorClient
}

type xxx_DefaultTraceDataCollectorClient struct {
	idatacollector.DataCollectorClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultTraceDataCollectorClient) DataCollector() idatacollector.DataCollectorClient {
	return o.DataCollectorClient
}

func (o *xxx_DefaultTraceDataCollectorClient) GetBufferSize(ctx context.Context, in *GetBufferSizeRequest, opts ...dcerpc.CallOption) (*GetBufferSizeResponse, error) {
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
	out := &GetBufferSizeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) SetBufferSize(ctx context.Context, in *SetBufferSizeRequest, opts ...dcerpc.CallOption) (*SetBufferSizeResponse, error) {
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
	out := &SetBufferSizeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetBuffersLost(ctx context.Context, in *GetBuffersLostRequest, opts ...dcerpc.CallOption) (*GetBuffersLostResponse, error) {
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
	out := &GetBuffersLostResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetBuffersWritten(ctx context.Context, in *GetBuffersWrittenRequest, opts ...dcerpc.CallOption) (*GetBuffersWrittenResponse, error) {
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
	out := &GetBuffersWrittenResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetClockType(ctx context.Context, in *GetClockTypeRequest, opts ...dcerpc.CallOption) (*GetClockTypeResponse, error) {
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
	out := &GetClockTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) SetClockType(ctx context.Context, in *SetClockTypeRequest, opts ...dcerpc.CallOption) (*SetClockTypeResponse, error) {
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
	out := &SetClockTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetEventsLost(ctx context.Context, in *GetEventsLostRequest, opts ...dcerpc.CallOption) (*GetEventsLostResponse, error) {
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
	out := &GetEventsLostResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetExtendedModes(ctx context.Context, in *GetExtendedModesRequest, opts ...dcerpc.CallOption) (*GetExtendedModesResponse, error) {
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
	out := &GetExtendedModesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) SetExtendedModes(ctx context.Context, in *SetExtendedModesRequest, opts ...dcerpc.CallOption) (*SetExtendedModesResponse, error) {
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
	out := &SetExtendedModesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetFlushTimer(ctx context.Context, in *GetFlushTimerRequest, opts ...dcerpc.CallOption) (*GetFlushTimerResponse, error) {
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
	out := &GetFlushTimerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) SetFlushTimer(ctx context.Context, in *SetFlushTimerRequest, opts ...dcerpc.CallOption) (*SetFlushTimerResponse, error) {
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
	out := &SetFlushTimerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetFreeBuffers(ctx context.Context, in *GetFreeBuffersRequest, opts ...dcerpc.CallOption) (*GetFreeBuffersResponse, error) {
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
	out := &GetFreeBuffersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetGUID(ctx context.Context, in *GetGUIDRequest, opts ...dcerpc.CallOption) (*GetGUIDResponse, error) {
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
	out := &GetGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) SetGUID(ctx context.Context, in *SetGUIDRequest, opts ...dcerpc.CallOption) (*SetGUIDResponse, error) {
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
	out := &SetGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetIsKernelTrace(ctx context.Context, in *GetIsKernelTraceRequest, opts ...dcerpc.CallOption) (*GetIsKernelTraceResponse, error) {
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
	out := &GetIsKernelTraceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetMaximumBuffers(ctx context.Context, in *GetMaximumBuffersRequest, opts ...dcerpc.CallOption) (*GetMaximumBuffersResponse, error) {
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
	out := &GetMaximumBuffersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) SetMaximumBuffers(ctx context.Context, in *SetMaximumBuffersRequest, opts ...dcerpc.CallOption) (*SetMaximumBuffersResponse, error) {
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
	out := &SetMaximumBuffersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetMinimumBuffers(ctx context.Context, in *GetMinimumBuffersRequest, opts ...dcerpc.CallOption) (*GetMinimumBuffersResponse, error) {
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
	out := &GetMinimumBuffersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) SetMinimumBuffers(ctx context.Context, in *SetMinimumBuffersRequest, opts ...dcerpc.CallOption) (*SetMinimumBuffersResponse, error) {
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
	out := &SetMinimumBuffersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetNumberOfBuffers(ctx context.Context, in *GetNumberOfBuffersRequest, opts ...dcerpc.CallOption) (*GetNumberOfBuffersResponse, error) {
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
	out := &GetNumberOfBuffersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) SetNumberOfBuffers(ctx context.Context, in *SetNumberOfBuffersRequest, opts ...dcerpc.CallOption) (*SetNumberOfBuffersResponse, error) {
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
	out := &SetNumberOfBuffersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetPreallocateFile(ctx context.Context, in *GetPreallocateFileRequest, opts ...dcerpc.CallOption) (*GetPreallocateFileResponse, error) {
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
	out := &GetPreallocateFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) SetPreallocateFile(ctx context.Context, in *SetPreallocateFileRequest, opts ...dcerpc.CallOption) (*SetPreallocateFileResponse, error) {
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
	out := &SetPreallocateFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetProcessMode(ctx context.Context, in *GetProcessModeRequest, opts ...dcerpc.CallOption) (*GetProcessModeResponse, error) {
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
	out := &GetProcessModeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) SetProcessMode(ctx context.Context, in *SetProcessModeRequest, opts ...dcerpc.CallOption) (*SetProcessModeResponse, error) {
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
	out := &SetProcessModeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetRealTimeBuffersLost(ctx context.Context, in *GetRealTimeBuffersLostRequest, opts ...dcerpc.CallOption) (*GetRealTimeBuffersLostResponse, error) {
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
	out := &GetRealTimeBuffersLostResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetSessionID(ctx context.Context, in *GetSessionIDRequest, opts ...dcerpc.CallOption) (*GetSessionIDResponse, error) {
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
	out := &GetSessionIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetSessionName(ctx context.Context, in *GetSessionNameRequest, opts ...dcerpc.CallOption) (*GetSessionNameResponse, error) {
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
	out := &GetSessionNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) SetSessionName(ctx context.Context, in *SetSessionNameRequest, opts ...dcerpc.CallOption) (*SetSessionNameResponse, error) {
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
	out := &SetSessionNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetSessionThreadID(ctx context.Context, in *GetSessionThreadIDRequest, opts ...dcerpc.CallOption) (*GetSessionThreadIDResponse, error) {
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
	out := &GetSessionThreadIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetStreamMode(ctx context.Context, in *GetStreamModeRequest, opts ...dcerpc.CallOption) (*GetStreamModeResponse, error) {
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
	out := &GetStreamModeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) SetStreamMode(ctx context.Context, in *SetStreamModeRequest, opts ...dcerpc.CallOption) (*SetStreamModeResponse, error) {
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
	out := &SetStreamModeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) GetTraceDataProviders(ctx context.Context, in *GetTraceDataProvidersRequest, opts ...dcerpc.CallOption) (*GetTraceDataProvidersResponse, error) {
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
	out := &GetTraceDataProvidersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTraceDataCollectorClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTraceDataCollectorClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultTraceDataCollectorClient) IPID(ctx context.Context, ipid *dcom.IPID) TraceDataCollectorClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultTraceDataCollectorClient{
		DataCollectorClient: o.DataCollectorClient.IPID(ctx, ipid),
		cc:                  o.cc,
		ipid:                ipid,
	}
}

func NewTraceDataCollectorClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TraceDataCollectorClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TraceDataCollectorSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idatacollector.NewDataCollectorClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultTraceDataCollectorClient{
		DataCollectorClient: base,
		cc:                  cc,
		ipid:                ipid,
	}, nil
}

// xxx_GetBufferSizeOperation structure represents the BufferSize operation
type xxx_GetBufferSizeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Size   uint32         `idl:"name:size" json:"size"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBufferSizeOperation) OpNum() int { return 32 }

func (o *xxx_GetBufferSizeOperation) OpName() string { return "/ITraceDataCollector/v0/BufferSize" }

func (o *xxx_GetBufferSizeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBufferSizeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetBufferSizeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetBufferSizeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBufferSizeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// size {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Size); err != nil {
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

func (o *xxx_GetBufferSizeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// size {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Size); err != nil {
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

// GetBufferSizeRequest structure represents the BufferSize operation request
type GetBufferSizeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetBufferSizeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetBufferSizeOperation) *xxx_GetBufferSizeOperation {
	if op == nil {
		op = &xxx_GetBufferSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetBufferSizeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetBufferSizeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetBufferSizeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetBufferSizeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBufferSizeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetBufferSizeResponse structure represents the BufferSize operation response
type GetBufferSizeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Size uint32         `idl:"name:size" json:"size"`
	// Return: The BufferSize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetBufferSizeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetBufferSizeOperation) *xxx_GetBufferSizeOperation {
	if op == nil {
		op = &xxx_GetBufferSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Size = o.Size
	op.Return = o.Return
	return op
}

func (o *GetBufferSizeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetBufferSizeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Size = op.Size
	o.Return = op.Return
}
func (o *GetBufferSizeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetBufferSizeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBufferSizeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetBufferSizeOperation structure represents the BufferSize operation
type xxx_SetBufferSizeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Size   uint32         `idl:"name:size" json:"size"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetBufferSizeOperation) OpNum() int { return 33 }

func (o *xxx_SetBufferSizeOperation) OpName() string { return "/ITraceDataCollector/v0/BufferSize" }

func (o *xxx_SetBufferSizeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBufferSizeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// size {in} (1:(uint32))
	{
		if err := w.WriteData(o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBufferSizeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// size {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBufferSizeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBufferSizeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetBufferSizeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetBufferSizeRequest structure represents the BufferSize operation request
type SetBufferSizeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Size uint32         `idl:"name:size" json:"size"`
}

func (o *SetBufferSizeRequest) xxx_ToOp(ctx context.Context, op *xxx_SetBufferSizeOperation) *xxx_SetBufferSizeOperation {
	if op == nil {
		op = &xxx_SetBufferSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Size = o.Size
	return op
}

func (o *SetBufferSizeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetBufferSizeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Size = op.Size
}
func (o *SetBufferSizeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetBufferSizeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetBufferSizeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetBufferSizeResponse structure represents the BufferSize operation response
type SetBufferSizeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The BufferSize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetBufferSizeResponse) xxx_ToOp(ctx context.Context, op *xxx_SetBufferSizeOperation) *xxx_SetBufferSizeOperation {
	if op == nil {
		op = &xxx_SetBufferSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetBufferSizeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetBufferSizeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetBufferSizeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetBufferSizeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetBufferSizeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetBuffersLostOperation structure represents the BuffersLost operation
type xxx_GetBuffersLostOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBuffersLostOperation) OpNum() int { return 34 }

func (o *xxx_GetBuffersLostOperation) OpName() string { return "/ITraceDataCollector/v0/BuffersLost" }

func (o *xxx_GetBuffersLostOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBuffersLostOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetBuffersLostOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetBuffersLostOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBuffersLostOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// buffers {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Buffers); err != nil {
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

func (o *xxx_GetBuffersLostOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// buffers {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Buffers); err != nil {
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

// GetBuffersLostRequest structure represents the BuffersLost operation request
type GetBuffersLostRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetBuffersLostRequest) xxx_ToOp(ctx context.Context, op *xxx_GetBuffersLostOperation) *xxx_GetBuffersLostOperation {
	if op == nil {
		op = &xxx_GetBuffersLostOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetBuffersLostRequest) xxx_FromOp(ctx context.Context, op *xxx_GetBuffersLostOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetBuffersLostRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetBuffersLostRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBuffersLostOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetBuffersLostResponse structure represents the BuffersLost operation response
type GetBuffersLostResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// buffers: Receives the number of buffers that had to be discarded. Specifies the number
	// of buffers that could not be written to the log file. The valid range is from 0x00000000
	// through 0xFFFFFFFF inclusive.
	Buffers uint32 `idl:"name:buffers" json:"buffers"`
	// Return: The BuffersLost return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetBuffersLostResponse) xxx_ToOp(ctx context.Context, op *xxx_GetBuffersLostOperation) *xxx_GetBuffersLostOperation {
	if op == nil {
		op = &xxx_GetBuffersLostOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Buffers = o.Buffers
	op.Return = o.Return
	return op
}

func (o *GetBuffersLostResponse) xxx_FromOp(ctx context.Context, op *xxx_GetBuffersLostOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Buffers = op.Buffers
	o.Return = op.Return
}
func (o *GetBuffersLostResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetBuffersLostResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBuffersLostOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetBuffersWrittenOperation structure represents the BuffersWritten operation
type xxx_GetBuffersWrittenOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBuffersWrittenOperation) OpNum() int { return 36 }

func (o *xxx_GetBuffersWrittenOperation) OpName() string {
	return "/ITraceDataCollector/v0/BuffersWritten"
}

func (o *xxx_GetBuffersWrittenOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBuffersWrittenOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetBuffersWrittenOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetBuffersWrittenOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBuffersWrittenOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// buffers {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Buffers); err != nil {
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

func (o *xxx_GetBuffersWrittenOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// buffers {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Buffers); err != nil {
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

// GetBuffersWrittenRequest structure represents the BuffersWritten operation request
type GetBuffersWrittenRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetBuffersWrittenRequest) xxx_ToOp(ctx context.Context, op *xxx_GetBuffersWrittenOperation) *xxx_GetBuffersWrittenOperation {
	if op == nil {
		op = &xxx_GetBuffersWrittenOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetBuffersWrittenRequest) xxx_FromOp(ctx context.Context, op *xxx_GetBuffersWrittenOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetBuffersWrittenRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetBuffersWrittenRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBuffersWrittenOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetBuffersWrittenResponse structure represents the BuffersWritten operation response
type GetBuffersWrittenResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// buffers: Receives the number of buffers accepted. The valid range is from 0x00000000
	// through 0xFFFFFFFF inclusive.
	Buffers uint32 `idl:"name:buffers" json:"buffers"`
	// Return: The BuffersWritten return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetBuffersWrittenResponse) xxx_ToOp(ctx context.Context, op *xxx_GetBuffersWrittenOperation) *xxx_GetBuffersWrittenOperation {
	if op == nil {
		op = &xxx_GetBuffersWrittenOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Buffers = o.Buffers
	op.Return = o.Return
	return op
}

func (o *GetBuffersWrittenResponse) xxx_FromOp(ctx context.Context, op *xxx_GetBuffersWrittenOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Buffers = op.Buffers
	o.Return = op.Return
}
func (o *GetBuffersWrittenResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetBuffersWrittenResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBuffersWrittenOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClockTypeOperation structure represents the ClockType operation
type xxx_GetClockTypeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Clock  pla.ClockType  `idl:"name:clock" json:"clock"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClockTypeOperation) OpNum() int { return 38 }

func (o *xxx_GetClockTypeOperation) OpName() string { return "/ITraceDataCollector/v0/ClockType" }

func (o *xxx_GetClockTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClockTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetClockTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetClockTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClockTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// clock {out, retval} (1:{pointer=ref}*(1))(2:{alias=ClockType}(enum))
	{
		if err := w.WriteEnum(uint16(o.Clock)); err != nil {
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

func (o *xxx_GetClockTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// clock {out, retval} (1:{pointer=ref}*(1))(2:{alias=ClockType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Clock)); err != nil {
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

// GetClockTypeRequest structure represents the ClockType operation request
type GetClockTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetClockTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetClockTypeOperation) *xxx_GetClockTypeOperation {
	if op == nil {
		op = &xxx_GetClockTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetClockTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClockTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetClockTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetClockTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClockTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClockTypeResponse structure represents the ClockType operation response
type GetClockTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Clock pla.ClockType  `idl:"name:clock" json:"clock"`
	// Return: The ClockType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClockTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetClockTypeOperation) *xxx_GetClockTypeOperation {
	if op == nil {
		op = &xxx_GetClockTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Clock = o.Clock
	op.Return = o.Return
	return op
}

func (o *GetClockTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClockTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Clock = op.Clock
	o.Return = op.Return
}
func (o *GetClockTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetClockTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClockTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetClockTypeOperation structure represents the ClockType operation
type xxx_SetClockTypeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Clock  pla.ClockType  `idl:"name:clock" json:"clock"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetClockTypeOperation) OpNum() int { return 39 }

func (o *xxx_SetClockTypeOperation) OpName() string { return "/ITraceDataCollector/v0/ClockType" }

func (o *xxx_SetClockTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClockTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// clock {in} (1:{alias=ClockType}(enum))
	{
		if err := w.WriteEnum(uint16(o.Clock)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClockTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// clock {in} (1:{alias=ClockType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Clock)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClockTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClockTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetClockTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetClockTypeRequest structure represents the ClockType operation request
type SetClockTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Clock pla.ClockType  `idl:"name:clock" json:"clock"`
}

func (o *SetClockTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_SetClockTypeOperation) *xxx_SetClockTypeOperation {
	if op == nil {
		op = &xxx_SetClockTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Clock = o.Clock
	return op
}

func (o *SetClockTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetClockTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Clock = op.Clock
}
func (o *SetClockTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetClockTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClockTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetClockTypeResponse structure represents the ClockType operation response
type SetClockTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ClockType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetClockTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_SetClockTypeOperation) *xxx_SetClockTypeOperation {
	if op == nil {
		op = &xxx_SetClockTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetClockTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetClockTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetClockTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetClockTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClockTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEventsLostOperation structure represents the EventsLost operation
type xxx_GetEventsLostOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Events uint32         `idl:"name:events" json:"events"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEventsLostOperation) OpNum() int { return 40 }

func (o *xxx_GetEventsLostOperation) OpName() string { return "/ITraceDataCollector/v0/EventsLost" }

func (o *xxx_GetEventsLostOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventsLostOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetEventsLostOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetEventsLostOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventsLostOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// events {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Events); err != nil {
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

func (o *xxx_GetEventsLostOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// events {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Events); err != nil {
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

// GetEventsLostRequest structure represents the EventsLost operation request
type GetEventsLostRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEventsLostRequest) xxx_ToOp(ctx context.Context, op *xxx_GetEventsLostOperation) *xxx_GetEventsLostOperation {
	if op == nil {
		op = &xxx_GetEventsLostOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetEventsLostRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEventsLostOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEventsLostRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetEventsLostRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventsLostOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEventsLostResponse structure represents the EventsLost operation response
type GetEventsLostResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// events: Receives the number of events that were not written to the buffer. Specifies
	// the number of events that were lost due to the lack of buffers to write to. The valid
	// range is from 0x00000000 through 0xFFFFFFFF inclusive.
	Events uint32 `idl:"name:events" json:"events"`
	// Return: The EventsLost return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEventsLostResponse) xxx_ToOp(ctx context.Context, op *xxx_GetEventsLostOperation) *xxx_GetEventsLostOperation {
	if op == nil {
		op = &xxx_GetEventsLostOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Events = o.Events
	op.Return = o.Return
	return op
}

func (o *GetEventsLostResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEventsLostOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Events = op.Events
	o.Return = op.Return
}
func (o *GetEventsLostResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetEventsLostResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventsLostOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetExtendedModesOperation structure represents the ExtendedModes operation
type xxx_GetExtendedModesOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Mode   uint32         `idl:"name:mode" json:"mode"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetExtendedModesOperation) OpNum() int { return 42 }

func (o *xxx_GetExtendedModesOperation) OpName() string {
	return "/ITraceDataCollector/v0/ExtendedModes"
}

func (o *xxx_GetExtendedModesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExtendedModesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetExtendedModesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetExtendedModesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExtendedModesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// mode {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Mode); err != nil {
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

func (o *xxx_GetExtendedModesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// mode {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Mode); err != nil {
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

// GetExtendedModesRequest structure represents the ExtendedModes operation request
type GetExtendedModesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetExtendedModesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetExtendedModesOperation) *xxx_GetExtendedModesOperation {
	if op == nil {
		op = &xxx_GetExtendedModesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetExtendedModesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetExtendedModesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetExtendedModesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetExtendedModesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExtendedModesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetExtendedModesResponse structure represents the ExtendedModes operation response
type GetExtendedModesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// mode: Supplies the log file modes not already set by other methods. The valid values
	// for ExtendedModes are specified in section 2.2.9.
	//
	// mode: Receives the log file mode not already set by this object using the other properties.
	// The valid values for ExtendedModes are specified in section 2.2.9.
	Mode uint32 `idl:"name:mode" json:"mode"`
	// Return: The ExtendedModes return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetExtendedModesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetExtendedModesOperation) *xxx_GetExtendedModesOperation {
	if op == nil {
		op = &xxx_GetExtendedModesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Mode = o.Mode
	op.Return = o.Return
	return op
}

func (o *GetExtendedModesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetExtendedModesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Mode = op.Mode
	o.Return = op.Return
}
func (o *GetExtendedModesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetExtendedModesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExtendedModesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetExtendedModesOperation structure represents the ExtendedModes operation
type xxx_SetExtendedModesOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Mode   uint32         `idl:"name:mode" json:"mode"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetExtendedModesOperation) OpNum() int { return 43 }

func (o *xxx_SetExtendedModesOperation) OpName() string {
	return "/ITraceDataCollector/v0/ExtendedModes"
}

func (o *xxx_SetExtendedModesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExtendedModesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// mode {in} (1:(uint32))
	{
		if err := w.WriteData(o.Mode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExtendedModesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// mode {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Mode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExtendedModesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExtendedModesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetExtendedModesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetExtendedModesRequest structure represents the ExtendedModes operation request
type SetExtendedModesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// mode: Supplies the log file modes not already set by other methods. The valid values
	// for ExtendedModes are specified in section 2.2.9.
	//
	// mode: Receives the log file mode not already set by this object using the other properties.
	// The valid values for ExtendedModes are specified in section 2.2.9.
	Mode uint32 `idl:"name:mode" json:"mode"`
}

func (o *SetExtendedModesRequest) xxx_ToOp(ctx context.Context, op *xxx_SetExtendedModesOperation) *xxx_SetExtendedModesOperation {
	if op == nil {
		op = &xxx_SetExtendedModesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Mode = o.Mode
	return op
}

func (o *SetExtendedModesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetExtendedModesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Mode = op.Mode
}
func (o *SetExtendedModesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetExtendedModesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExtendedModesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetExtendedModesResponse structure represents the ExtendedModes operation response
type SetExtendedModesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ExtendedModes return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetExtendedModesResponse) xxx_ToOp(ctx context.Context, op *xxx_SetExtendedModesOperation) *xxx_SetExtendedModesOperation {
	if op == nil {
		op = &xxx_SetExtendedModesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetExtendedModesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetExtendedModesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetExtendedModesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetExtendedModesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExtendedModesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFlushTimerOperation structure represents the FlushTimer operation
type xxx_GetFlushTimerOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Seconds uint32         `idl:"name:seconds" json:"seconds"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFlushTimerOperation) OpNum() int { return 44 }

func (o *xxx_GetFlushTimerOperation) OpName() string { return "/ITraceDataCollector/v0/FlushTimer" }

func (o *xxx_GetFlushTimerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFlushTimerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFlushTimerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFlushTimerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFlushTimerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// seconds {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Seconds); err != nil {
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

func (o *xxx_GetFlushTimerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// seconds {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Seconds); err != nil {
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

// GetFlushTimerRequest structure represents the FlushTimer operation request
type GetFlushTimerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFlushTimerRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFlushTimerOperation) *xxx_GetFlushTimerOperation {
	if op == nil {
		op = &xxx_GetFlushTimerOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetFlushTimerRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFlushTimerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFlushTimerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFlushTimerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFlushTimerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFlushTimerResponse structure represents the FlushTimer operation response
type GetFlushTimerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Seconds uint32         `idl:"name:seconds" json:"seconds"`
	// Return: The FlushTimer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFlushTimerResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFlushTimerOperation) *xxx_GetFlushTimerOperation {
	if op == nil {
		op = &xxx_GetFlushTimerOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Seconds = o.Seconds
	op.Return = o.Return
	return op
}

func (o *GetFlushTimerResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFlushTimerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Seconds = op.Seconds
	o.Return = op.Return
}
func (o *GetFlushTimerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFlushTimerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFlushTimerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFlushTimerOperation structure represents the FlushTimer operation
type xxx_SetFlushTimerOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Seconds uint32         `idl:"name:seconds" json:"seconds"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFlushTimerOperation) OpNum() int { return 45 }

func (o *xxx_SetFlushTimerOperation) OpName() string { return "/ITraceDataCollector/v0/FlushTimer" }

func (o *xxx_SetFlushTimerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlushTimerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// seconds {in} (1:(uint32))
	{
		if err := w.WriteData(o.Seconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlushTimerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// seconds {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Seconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlushTimerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlushTimerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFlushTimerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFlushTimerRequest structure represents the FlushTimer operation request
type SetFlushTimerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Seconds uint32         `idl:"name:seconds" json:"seconds"`
}

func (o *SetFlushTimerRequest) xxx_ToOp(ctx context.Context, op *xxx_SetFlushTimerOperation) *xxx_SetFlushTimerOperation {
	if op == nil {
		op = &xxx_SetFlushTimerOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Seconds = o.Seconds
	return op
}

func (o *SetFlushTimerRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFlushTimerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Seconds = op.Seconds
}
func (o *SetFlushTimerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetFlushTimerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFlushTimerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFlushTimerResponse structure represents the FlushTimer operation response
type SetFlushTimerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The FlushTimer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFlushTimerResponse) xxx_ToOp(ctx context.Context, op *xxx_SetFlushTimerOperation) *xxx_SetFlushTimerOperation {
	if op == nil {
		op = &xxx_SetFlushTimerOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetFlushTimerResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFlushTimerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFlushTimerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetFlushTimerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFlushTimerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFreeBuffersOperation structure represents the FreeBuffers operation
type xxx_GetFreeBuffersOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFreeBuffersOperation) OpNum() int { return 46 }

func (o *xxx_GetFreeBuffersOperation) OpName() string { return "/ITraceDataCollector/v0/FreeBuffers" }

func (o *xxx_GetFreeBuffersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFreeBuffersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFreeBuffersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFreeBuffersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFreeBuffersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// buffers {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Buffers); err != nil {
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

func (o *xxx_GetFreeBuffersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// buffers {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Buffers); err != nil {
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

// GetFreeBuffersRequest structure represents the FreeBuffers operation request
type GetFreeBuffersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFreeBuffersRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFreeBuffersOperation) *xxx_GetFreeBuffersOperation {
	if op == nil {
		op = &xxx_GetFreeBuffersOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetFreeBuffersRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFreeBuffersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFreeBuffersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFreeBuffersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFreeBuffersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFreeBuffersResponse structure represents the FreeBuffers operation response
type GetFreeBuffersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// buffers: Receives the number of buffers that are allocated but unused in the event
	// tracing session's buffer pool. The valid range is from 0x00000000 through 0xFFFFFFFF
	// inclusive.
	Buffers uint32 `idl:"name:buffers" json:"buffers"`
	// Return: The FreeBuffers return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFreeBuffersResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFreeBuffersOperation) *xxx_GetFreeBuffersOperation {
	if op == nil {
		op = &xxx_GetFreeBuffersOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Buffers = o.Buffers
	op.Return = o.Return
	return op
}

func (o *GetFreeBuffersResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFreeBuffersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Buffers = op.Buffers
	o.Return = op.Return
}
func (o *GetFreeBuffersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFreeBuffersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFreeBuffersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetGUIDOperation structure represents the Guid operation
type xxx_GetGUIDOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	GUID   *dtyp.GUID     `idl:"name:guid" json:"guid"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetGUIDOperation) OpNum() int { return 48 }

func (o *xxx_GetGUIDOperation) OpName() string { return "/ITraceDataCollector/v0/Guid" }

func (o *xxx_GetGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// guid {out, retval} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
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

func (o *xxx_GetGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// guid {out, retval} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
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

// GetGUIDRequest structure represents the Guid operation request
type GetGUIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetGUIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetGUIDOperation) *xxx_GetGUIDOperation {
	if op == nil {
		op = &xxx_GetGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetGUIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetGUIDResponse structure represents the Guid operation response
type GetGUIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	GUID *dtyp.GUID     `idl:"name:guid" json:"guid"`
	// Return: The Guid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetGUIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetGUIDOperation) *xxx_GetGUIDOperation {
	if op == nil {
		op = &xxx_GetGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.GUID = o.GUID
	op.Return = o.Return
	return op
}

func (o *GetGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetGUIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.GUID = op.GUID
	o.Return = op.Return
}
func (o *GetGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetGUIDOperation structure represents the Guid operation
type xxx_SetGUIDOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	GUID   *dtyp.GUID     `idl:"name:guid" json:"guid"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetGUIDOperation) OpNum() int { return 49 }

func (o *xxx_SetGUIDOperation) OpName() string { return "/ITraceDataCollector/v0/Guid" }

func (o *xxx_SetGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// guid {in} (1:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// guid {in} (1:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetGUIDRequest structure represents the Guid operation request
type SetGUIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	GUID *dtyp.GUID     `idl:"name:guid" json:"guid"`
}

func (o *SetGUIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetGUIDOperation) *xxx_SetGUIDOperation {
	if op == nil {
		op = &xxx_SetGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.GUID = o.GUID
	return op
}

func (o *SetGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetGUIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.GUID = op.GUID
}
func (o *SetGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetGUIDResponse structure represents the Guid operation response
type SetGUIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Guid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetGUIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetGUIDOperation) *xxx_SetGUIDOperation {
	if op == nil {
		op = &xxx_SetGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetGUIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsKernelTraceOperation structure represents the IsKernelTrace operation
type xxx_GetIsKernelTraceOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Kernel int16          `idl:"name:kernel" json:"kernel"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsKernelTraceOperation) OpNum() int { return 50 }

func (o *xxx_GetIsKernelTraceOperation) OpName() string {
	return "/ITraceDataCollector/v0/IsKernelTrace"
}

func (o *xxx_GetIsKernelTraceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsKernelTraceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsKernelTraceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsKernelTraceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsKernelTraceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// kernel {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Kernel); err != nil {
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

func (o *xxx_GetIsKernelTraceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// kernel {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Kernel); err != nil {
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

// GetIsKernelTraceRequest structure represents the IsKernelTrace operation request
type GetIsKernelTraceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsKernelTraceRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIsKernelTraceOperation) *xxx_GetIsKernelTraceOperation {
	if op == nil {
		op = &xxx_GetIsKernelTraceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsKernelTraceRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsKernelTraceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsKernelTraceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsKernelTraceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsKernelTraceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsKernelTraceResponse structure represents the IsKernelTrace operation response
type GetIsKernelTraceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// kernel: Receives VARIANT_TRUE if the trace contains kernel providers.
	Kernel int16 `idl:"name:kernel" json:"kernel"`
	// Return: The IsKernelTrace return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsKernelTraceResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIsKernelTraceOperation) *xxx_GetIsKernelTraceOperation {
	if op == nil {
		op = &xxx_GetIsKernelTraceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Kernel = o.Kernel
	op.Return = o.Return
	return op
}

func (o *GetIsKernelTraceResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsKernelTraceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Kernel = op.Kernel
	o.Return = op.Return
}
func (o *GetIsKernelTraceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsKernelTraceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsKernelTraceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMaximumBuffersOperation structure represents the MaximumBuffers operation
type xxx_GetMaximumBuffersOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMaximumBuffersOperation) OpNum() int { return 51 }

func (o *xxx_GetMaximumBuffersOperation) OpName() string {
	return "/ITraceDataCollector/v0/MaximumBuffers"
}

func (o *xxx_GetMaximumBuffersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMaximumBuffersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMaximumBuffersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMaximumBuffersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMaximumBuffersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// buffers {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Buffers); err != nil {
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

func (o *xxx_GetMaximumBuffersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// buffers {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Buffers); err != nil {
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

// GetMaximumBuffersRequest structure represents the MaximumBuffers operation request
type GetMaximumBuffersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMaximumBuffersRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMaximumBuffersOperation) *xxx_GetMaximumBuffersOperation {
	if op == nil {
		op = &xxx_GetMaximumBuffersOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMaximumBuffersRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMaximumBuffersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMaximumBuffersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMaximumBuffersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMaximumBuffersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMaximumBuffersResponse structure represents the MaximumBuffers operation response
type GetMaximumBuffersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
	// Return: The MaximumBuffers return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMaximumBuffersResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMaximumBuffersOperation) *xxx_GetMaximumBuffersOperation {
	if op == nil {
		op = &xxx_GetMaximumBuffersOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Buffers = o.Buffers
	op.Return = o.Return
	return op
}

func (o *GetMaximumBuffersResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMaximumBuffersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Buffers = op.Buffers
	o.Return = op.Return
}
func (o *GetMaximumBuffersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMaximumBuffersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMaximumBuffersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMaximumBuffersOperation structure represents the MaximumBuffers operation
type xxx_SetMaximumBuffersOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMaximumBuffersOperation) OpNum() int { return 52 }

func (o *xxx_SetMaximumBuffersOperation) OpName() string {
	return "/ITraceDataCollector/v0/MaximumBuffers"
}

func (o *xxx_SetMaximumBuffersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaximumBuffersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// buffers {in} (1:(uint32))
	{
		if err := w.WriteData(o.Buffers); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaximumBuffersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// buffers {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Buffers); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaximumBuffersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaximumBuffersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMaximumBuffersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMaximumBuffersRequest structure represents the MaximumBuffers operation request
type SetMaximumBuffersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
}

func (o *SetMaximumBuffersRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMaximumBuffersOperation) *xxx_SetMaximumBuffersOperation {
	if op == nil {
		op = &xxx_SetMaximumBuffersOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Buffers = o.Buffers
	return op
}

func (o *SetMaximumBuffersRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMaximumBuffersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Buffers = op.Buffers
}
func (o *SetMaximumBuffersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMaximumBuffersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMaximumBuffersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMaximumBuffersResponse structure represents the MaximumBuffers operation response
type SetMaximumBuffersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MaximumBuffers return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMaximumBuffersResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMaximumBuffersOperation) *xxx_SetMaximumBuffersOperation {
	if op == nil {
		op = &xxx_SetMaximumBuffersOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMaximumBuffersResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMaximumBuffersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMaximumBuffersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMaximumBuffersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMaximumBuffersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMinimumBuffersOperation structure represents the MinimumBuffers operation
type xxx_GetMinimumBuffersOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMinimumBuffersOperation) OpNum() int { return 53 }

func (o *xxx_GetMinimumBuffersOperation) OpName() string {
	return "/ITraceDataCollector/v0/MinimumBuffers"
}

func (o *xxx_GetMinimumBuffersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMinimumBuffersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMinimumBuffersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMinimumBuffersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMinimumBuffersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// buffers {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Buffers); err != nil {
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

func (o *xxx_GetMinimumBuffersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// buffers {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Buffers); err != nil {
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

// GetMinimumBuffersRequest structure represents the MinimumBuffers operation request
type GetMinimumBuffersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMinimumBuffersRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMinimumBuffersOperation) *xxx_GetMinimumBuffersOperation {
	if op == nil {
		op = &xxx_GetMinimumBuffersOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMinimumBuffersRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMinimumBuffersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMinimumBuffersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMinimumBuffersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMinimumBuffersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMinimumBuffersResponse structure represents the MinimumBuffers operation response
type GetMinimumBuffersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
	// Return: The MinimumBuffers return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMinimumBuffersResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMinimumBuffersOperation) *xxx_GetMinimumBuffersOperation {
	if op == nil {
		op = &xxx_GetMinimumBuffersOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Buffers = o.Buffers
	op.Return = o.Return
	return op
}

func (o *GetMinimumBuffersResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMinimumBuffersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Buffers = op.Buffers
	o.Return = op.Return
}
func (o *GetMinimumBuffersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMinimumBuffersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMinimumBuffersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMinimumBuffersOperation structure represents the MinimumBuffers operation
type xxx_SetMinimumBuffersOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMinimumBuffersOperation) OpNum() int { return 54 }

func (o *xxx_SetMinimumBuffersOperation) OpName() string {
	return "/ITraceDataCollector/v0/MinimumBuffers"
}

func (o *xxx_SetMinimumBuffersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMinimumBuffersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// buffers {in} (1:(uint32))
	{
		if err := w.WriteData(o.Buffers); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMinimumBuffersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// buffers {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Buffers); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMinimumBuffersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMinimumBuffersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMinimumBuffersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMinimumBuffersRequest structure represents the MinimumBuffers operation request
type SetMinimumBuffersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
}

func (o *SetMinimumBuffersRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMinimumBuffersOperation) *xxx_SetMinimumBuffersOperation {
	if op == nil {
		op = &xxx_SetMinimumBuffersOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Buffers = o.Buffers
	return op
}

func (o *SetMinimumBuffersRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMinimumBuffersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Buffers = op.Buffers
}
func (o *SetMinimumBuffersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMinimumBuffersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMinimumBuffersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMinimumBuffersResponse structure represents the MinimumBuffers operation response
type SetMinimumBuffersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MinimumBuffers return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMinimumBuffersResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMinimumBuffersOperation) *xxx_SetMinimumBuffersOperation {
	if op == nil {
		op = &xxx_SetMinimumBuffersOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMinimumBuffersResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMinimumBuffersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMinimumBuffersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMinimumBuffersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMinimumBuffersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNumberOfBuffersOperation structure represents the NumberOfBuffers operation
type xxx_GetNumberOfBuffersOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNumberOfBuffersOperation) OpNum() int { return 55 }

func (o *xxx_GetNumberOfBuffersOperation) OpName() string {
	return "/ITraceDataCollector/v0/NumberOfBuffers"
}

func (o *xxx_GetNumberOfBuffersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNumberOfBuffersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetNumberOfBuffersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetNumberOfBuffersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNumberOfBuffersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// buffers {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Buffers); err != nil {
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

func (o *xxx_GetNumberOfBuffersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// buffers {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Buffers); err != nil {
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

// GetNumberOfBuffersRequest structure represents the NumberOfBuffers operation request
type GetNumberOfBuffersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetNumberOfBuffersRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNumberOfBuffersOperation) *xxx_GetNumberOfBuffersOperation {
	if op == nil {
		op = &xxx_GetNumberOfBuffersOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetNumberOfBuffersRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNumberOfBuffersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetNumberOfBuffersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNumberOfBuffersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNumberOfBuffersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNumberOfBuffersResponse structure represents the NumberOfBuffers operation response
type GetNumberOfBuffersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
	// Return: The NumberOfBuffers return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNumberOfBuffersResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNumberOfBuffersOperation) *xxx_GetNumberOfBuffersOperation {
	if op == nil {
		op = &xxx_GetNumberOfBuffersOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Buffers = o.Buffers
	op.Return = o.Return
	return op
}

func (o *GetNumberOfBuffersResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNumberOfBuffersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Buffers = op.Buffers
	o.Return = op.Return
}
func (o *GetNumberOfBuffersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNumberOfBuffersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNumberOfBuffersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNumberOfBuffersOperation structure represents the NumberOfBuffers operation
type xxx_SetNumberOfBuffersOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNumberOfBuffersOperation) OpNum() int { return 56 }

func (o *xxx_SetNumberOfBuffersOperation) OpName() string {
	return "/ITraceDataCollector/v0/NumberOfBuffers"
}

func (o *xxx_SetNumberOfBuffersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNumberOfBuffersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// buffers {in} (1:(uint32))
	{
		if err := w.WriteData(o.Buffers); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNumberOfBuffersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// buffers {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Buffers); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNumberOfBuffersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNumberOfBuffersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNumberOfBuffersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNumberOfBuffersRequest structure represents the NumberOfBuffers operation request
type SetNumberOfBuffersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
}

func (o *SetNumberOfBuffersRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNumberOfBuffersOperation) *xxx_SetNumberOfBuffersOperation {
	if op == nil {
		op = &xxx_SetNumberOfBuffersOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Buffers = o.Buffers
	return op
}

func (o *SetNumberOfBuffersRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNumberOfBuffersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Buffers = op.Buffers
}
func (o *SetNumberOfBuffersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNumberOfBuffersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNumberOfBuffersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNumberOfBuffersResponse structure represents the NumberOfBuffers operation response
type SetNumberOfBuffersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The NumberOfBuffers return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNumberOfBuffersResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNumberOfBuffersOperation) *xxx_SetNumberOfBuffersOperation {
	if op == nil {
		op = &xxx_SetNumberOfBuffersOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNumberOfBuffersResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNumberOfBuffersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNumberOfBuffersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNumberOfBuffersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNumberOfBuffersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPreallocateFileOperation structure represents the PreallocateFile operation
type xxx_GetPreallocateFileOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Allocate int16          `idl:"name:allocate" json:"allocate"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPreallocateFileOperation) OpNum() int { return 57 }

func (o *xxx_GetPreallocateFileOperation) OpName() string {
	return "/ITraceDataCollector/v0/PreallocateFile"
}

func (o *xxx_GetPreallocateFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPreallocateFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPreallocateFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPreallocateFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPreallocateFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// allocate {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Allocate); err != nil {
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

func (o *xxx_GetPreallocateFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// allocate {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Allocate); err != nil {
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

// GetPreallocateFileRequest structure represents the PreallocateFile operation request
type GetPreallocateFileRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPreallocateFileRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPreallocateFileOperation) *xxx_GetPreallocateFileOperation {
	if op == nil {
		op = &xxx_GetPreallocateFileOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPreallocateFileRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPreallocateFileOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPreallocateFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPreallocateFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPreallocateFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPreallocateFileResponse structure represents the PreallocateFile operation response
type GetPreallocateFileResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Allocate int16          `idl:"name:allocate" json:"allocate"`
	// Return: The PreallocateFile return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPreallocateFileResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPreallocateFileOperation) *xxx_GetPreallocateFileOperation {
	if op == nil {
		op = &xxx_GetPreallocateFileOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Allocate = o.Allocate
	op.Return = o.Return
	return op
}

func (o *GetPreallocateFileResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPreallocateFileOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Allocate = op.Allocate
	o.Return = op.Return
}
func (o *GetPreallocateFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPreallocateFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPreallocateFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPreallocateFileOperation structure represents the PreallocateFile operation
type xxx_SetPreallocateFileOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Allocate int16          `idl:"name:allocate" json:"allocate"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPreallocateFileOperation) OpNum() int { return 58 }

func (o *xxx_SetPreallocateFileOperation) OpName() string {
	return "/ITraceDataCollector/v0/PreallocateFile"
}

func (o *xxx_SetPreallocateFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPreallocateFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// allocate {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Allocate); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPreallocateFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// allocate {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Allocate); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPreallocateFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPreallocateFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetPreallocateFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetPreallocateFileRequest structure represents the PreallocateFile operation request
type SetPreallocateFileRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	Allocate int16          `idl:"name:allocate" json:"allocate"`
}

func (o *SetPreallocateFileRequest) xxx_ToOp(ctx context.Context, op *xxx_SetPreallocateFileOperation) *xxx_SetPreallocateFileOperation {
	if op == nil {
		op = &xxx_SetPreallocateFileOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Allocate = o.Allocate
	return op
}

func (o *SetPreallocateFileRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPreallocateFileOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Allocate = op.Allocate
}
func (o *SetPreallocateFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetPreallocateFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPreallocateFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPreallocateFileResponse structure represents the PreallocateFile operation response
type SetPreallocateFileResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PreallocateFile return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPreallocateFileResponse) xxx_ToOp(ctx context.Context, op *xxx_SetPreallocateFileOperation) *xxx_SetPreallocateFileOperation {
	if op == nil {
		op = &xxx_SetPreallocateFileOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetPreallocateFileResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPreallocateFileOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetPreallocateFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetPreallocateFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPreallocateFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetProcessModeOperation structure represents the ProcessMode operation
type xxx_GetProcessModeOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Process int16          `idl:"name:process" json:"process"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetProcessModeOperation) OpNum() int { return 59 }

func (o *xxx_GetProcessModeOperation) OpName() string { return "/ITraceDataCollector/v0/ProcessMode" }

func (o *xxx_GetProcessModeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProcessModeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetProcessModeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetProcessModeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProcessModeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// process {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Process); err != nil {
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

func (o *xxx_GetProcessModeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// process {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Process); err != nil {
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

// GetProcessModeRequest structure represents the ProcessMode operation request
type GetProcessModeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetProcessModeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetProcessModeOperation) *xxx_GetProcessModeOperation {
	if op == nil {
		op = &xxx_GetProcessModeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetProcessModeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetProcessModeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetProcessModeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetProcessModeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProcessModeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetProcessModeResponse structure represents the ProcessMode operation response
type GetProcessModeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Process int16          `idl:"name:process" json:"process"`
	// Return: The ProcessMode return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetProcessModeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetProcessModeOperation) *xxx_GetProcessModeOperation {
	if op == nil {
		op = &xxx_GetProcessModeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Process = o.Process
	op.Return = o.Return
	return op
}

func (o *GetProcessModeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetProcessModeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Process = op.Process
	o.Return = op.Return
}
func (o *GetProcessModeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetProcessModeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProcessModeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetProcessModeOperation structure represents the ProcessMode operation
type xxx_SetProcessModeOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Process int16          `idl:"name:process" json:"process"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetProcessModeOperation) OpNum() int { return 60 }

func (o *xxx_SetProcessModeOperation) OpName() string { return "/ITraceDataCollector/v0/ProcessMode" }

func (o *xxx_SetProcessModeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetProcessModeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// process {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Process); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetProcessModeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// process {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Process); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetProcessModeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetProcessModeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetProcessModeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetProcessModeRequest structure represents the ProcessMode operation request
type SetProcessModeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Process int16          `idl:"name:process" json:"process"`
}

func (o *SetProcessModeRequest) xxx_ToOp(ctx context.Context, op *xxx_SetProcessModeOperation) *xxx_SetProcessModeOperation {
	if op == nil {
		op = &xxx_SetProcessModeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Process = o.Process
	return op
}

func (o *SetProcessModeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetProcessModeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Process = op.Process
}
func (o *SetProcessModeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetProcessModeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetProcessModeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetProcessModeResponse structure represents the ProcessMode operation response
type SetProcessModeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ProcessMode return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetProcessModeResponse) xxx_ToOp(ctx context.Context, op *xxx_SetProcessModeOperation) *xxx_SetProcessModeOperation {
	if op == nil {
		op = &xxx_SetProcessModeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetProcessModeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetProcessModeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetProcessModeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetProcessModeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetProcessModeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRealTimeBuffersLostOperation structure represents the RealTimeBuffersLost operation
type xxx_GetRealTimeBuffersLostOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Buffers uint32         `idl:"name:buffers" json:"buffers"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRealTimeBuffersLostOperation) OpNum() int { return 61 }

func (o *xxx_GetRealTimeBuffersLostOperation) OpName() string {
	return "/ITraceDataCollector/v0/RealTimeBuffersLost"
}

func (o *xxx_GetRealTimeBuffersLostOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRealTimeBuffersLostOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetRealTimeBuffersLostOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetRealTimeBuffersLostOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRealTimeBuffersLostOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// buffers {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Buffers); err != nil {
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

func (o *xxx_GetRealTimeBuffersLostOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// buffers {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Buffers); err != nil {
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

// GetRealTimeBuffersLostRequest structure represents the RealTimeBuffersLost operation request
type GetRealTimeBuffersLostRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetRealTimeBuffersLostRequest) xxx_ToOp(ctx context.Context, op *xxx_GetRealTimeBuffersLostOperation) *xxx_GetRealTimeBuffersLostOperation {
	if op == nil {
		op = &xxx_GetRealTimeBuffersLostOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetRealTimeBuffersLostRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRealTimeBuffersLostOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetRealTimeBuffersLostRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetRealTimeBuffersLostRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRealTimeBuffersLostOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRealTimeBuffersLostResponse structure represents the RealTimeBuffersLost operation response
type GetRealTimeBuffersLostResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// buffers: The number of buffers that could not be delivered in real-time. The valid
	// range is from 0x00000000 through 0xFFFFFFFF inclusive.
	Buffers uint32 `idl:"name:buffers" json:"buffers"`
	// Return: The RealTimeBuffersLost return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRealTimeBuffersLostResponse) xxx_ToOp(ctx context.Context, op *xxx_GetRealTimeBuffersLostOperation) *xxx_GetRealTimeBuffersLostOperation {
	if op == nil {
		op = &xxx_GetRealTimeBuffersLostOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Buffers = o.Buffers
	op.Return = o.Return
	return op
}

func (o *GetRealTimeBuffersLostResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRealTimeBuffersLostOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Buffers = op.Buffers
	o.Return = op.Return
}
func (o *GetRealTimeBuffersLostResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetRealTimeBuffersLostResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRealTimeBuffersLostOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSessionIDOperation structure represents the SessionId operation
type xxx_GetSessionIDOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	ID     uint64         `idl:"name:id" json:"id"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSessionIDOperation) OpNum() int { return 63 }

func (o *xxx_GetSessionIDOperation) OpName() string { return "/ITraceDataCollector/v0/SessionId" }

func (o *xxx_GetSessionIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSessionIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSessionIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// id {out, retval} (1:{pointer=ref}*(1))(2:{alias=ULONG64}(uint64))
	{
		if err := w.WriteData(o.ID); err != nil {
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

func (o *xxx_GetSessionIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// id {out, retval} (1:{pointer=ref}*(1))(2:{alias=ULONG64}(uint64))
	{
		if err := w.ReadData(&o.ID); err != nil {
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

// GetSessionIDRequest structure represents the SessionId operation request
type GetSessionIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSessionIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSessionIDOperation) *xxx_GetSessionIDOperation {
	if op == nil {
		op = &xxx_GetSessionIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSessionIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSessionIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSessionIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSessionIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSessionIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSessionIDResponse structure represents the SessionId operation response
type GetSessionIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// id: Receives the session identifier. Only the lower 2 bytes of the id value are specified;
	// therefore, the valid range of these lower 2 bytes is from 0x0000 to 0x003F. The upper
	// 6 bytes of the id MUST be ignored.
	ID uint64 `idl:"name:id" json:"id"`
	// Return: The SessionId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSessionIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSessionIDOperation) *xxx_GetSessionIDOperation {
	if op == nil {
		op = &xxx_GetSessionIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ID = o.ID
	op.Return = o.Return
	return op
}

func (o *GetSessionIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSessionIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ID = op.ID
	o.Return = op.Return
}
func (o *GetSessionIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSessionIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSessionIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSessionNameOperation structure represents the SessionName operation
type xxx_GetSessionNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSessionNameOperation) OpNum() int { return 65 }

func (o *xxx_GetSessionNameOperation) OpName() string { return "/ITraceDataCollector/v0/SessionName" }

func (o *xxx_GetSessionNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSessionNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSessionNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSessionNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetSessionNameRequest structure represents the SessionName operation request
type GetSessionNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSessionNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSessionNameOperation) *xxx_GetSessionNameOperation {
	if op == nil {
		op = &xxx_GetSessionNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSessionNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSessionNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSessionNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSessionNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSessionNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSessionNameResponse structure represents the SessionName operation response
type GetSessionNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name *oaut.String   `idl:"name:name" json:"name"`
	// Return: The SessionName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSessionNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSessionNameOperation) *xxx_GetSessionNameOperation {
	if op == nil {
		op = &xxx_GetSessionNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Name = o.Name
	op.Return = o.Return
	return op
}

func (o *GetSessionNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSessionNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Name = op.Name
	o.Return = op.Return
}
func (o *GetSessionNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSessionNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSessionNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSessionNameOperation structure represents the SessionName operation
type xxx_SetSessionNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSessionNameOperation) OpNum() int { return 66 }

func (o *xxx_SetSessionNameOperation) OpName() string { return "/ITraceDataCollector/v0/SessionName" }

func (o *xxx_SetSessionNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSessionNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSessionNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_SetSessionNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSessionNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSessionNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSessionNameRequest structure represents the SessionName operation request
type SetSessionNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Name *oaut.String   `idl:"name:name" json:"name"`
}

func (o *SetSessionNameRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSessionNameOperation) *xxx_SetSessionNameOperation {
	if op == nil {
		op = &xxx_SetSessionNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Name = o.Name
	return op
}

func (o *SetSessionNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSessionNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
}
func (o *SetSessionNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSessionNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSessionNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSessionNameResponse structure represents the SessionName operation response
type SetSessionNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SessionName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSessionNameResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSessionNameOperation) *xxx_SetSessionNameOperation {
	if op == nil {
		op = &xxx_SetSessionNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSessionNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSessionNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSessionNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSessionNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSessionNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSessionThreadIDOperation structure represents the SessionThreadId operation
type xxx_GetSessionThreadIDOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	TID    uint32         `idl:"name:tid" json:"tid"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSessionThreadIDOperation) OpNum() int { return 67 }

func (o *xxx_GetSessionThreadIDOperation) OpName() string {
	return "/ITraceDataCollector/v0/SessionThreadId"
}

func (o *xxx_GetSessionThreadIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionThreadIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSessionThreadIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSessionThreadIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionThreadIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// tid {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TID); err != nil {
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

func (o *xxx_GetSessionThreadIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// tid {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TID); err != nil {
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

// GetSessionThreadIDRequest structure represents the SessionThreadId operation request
type GetSessionThreadIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSessionThreadIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSessionThreadIDOperation) *xxx_GetSessionThreadIDOperation {
	if op == nil {
		op = &xxx_GetSessionThreadIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSessionThreadIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSessionThreadIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSessionThreadIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSessionThreadIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSessionThreadIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSessionThreadIDResponse structure represents the SessionThreadId operation response
type GetSessionThreadIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// tid: Receives the current thread of the log session, if running.
	TID uint32 `idl:"name:tid" json:"tid"`
	// Return: The SessionThreadId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSessionThreadIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSessionThreadIDOperation) *xxx_GetSessionThreadIDOperation {
	if op == nil {
		op = &xxx_GetSessionThreadIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.TID = o.TID
	op.Return = o.Return
	return op
}

func (o *GetSessionThreadIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSessionThreadIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TID = op.TID
	o.Return = op.Return
}
func (o *GetSessionThreadIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSessionThreadIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSessionThreadIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetStreamModeOperation structure represents the StreamMode operation
type xxx_GetStreamModeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Mode   pla.StreamMode `idl:"name:mode" json:"mode"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetStreamModeOperation) OpNum() int { return 69 }

func (o *xxx_GetStreamModeOperation) OpName() string { return "/ITraceDataCollector/v0/StreamMode" }

func (o *xxx_GetStreamModeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStreamModeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetStreamModeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetStreamModeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStreamModeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// mode {out, retval} (1:{pointer=ref}*(1))(2:{alias=StreamMode}(enum))
	{
		if err := w.WriteEnum(uint16(o.Mode)); err != nil {
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

func (o *xxx_GetStreamModeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// mode {out, retval} (1:{pointer=ref}*(1))(2:{alias=StreamMode}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Mode)); err != nil {
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

// GetStreamModeRequest structure represents the StreamMode operation request
type GetStreamModeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetStreamModeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetStreamModeOperation) *xxx_GetStreamModeOperation {
	if op == nil {
		op = &xxx_GetStreamModeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetStreamModeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetStreamModeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetStreamModeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetStreamModeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStreamModeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetStreamModeResponse structure represents the StreamMode operation response
type GetStreamModeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Mode pla.StreamMode `idl:"name:mode" json:"mode"`
	// Return: The StreamMode return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetStreamModeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetStreamModeOperation) *xxx_GetStreamModeOperation {
	if op == nil {
		op = &xxx_GetStreamModeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Mode = o.Mode
	op.Return = o.Return
	return op
}

func (o *GetStreamModeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetStreamModeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Mode = op.Mode
	o.Return = op.Return
}
func (o *GetStreamModeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetStreamModeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStreamModeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetStreamModeOperation structure represents the StreamMode operation
type xxx_SetStreamModeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Mode   pla.StreamMode `idl:"name:mode" json:"mode"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetStreamModeOperation) OpNum() int { return 70 }

func (o *xxx_SetStreamModeOperation) OpName() string { return "/ITraceDataCollector/v0/StreamMode" }

func (o *xxx_SetStreamModeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStreamModeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// mode {in} (1:{alias=StreamMode}(enum))
	{
		if err := w.WriteEnum(uint16(o.Mode)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStreamModeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// mode {in} (1:{alias=StreamMode}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Mode)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStreamModeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStreamModeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetStreamModeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetStreamModeRequest structure represents the StreamMode operation request
type SetStreamModeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Mode pla.StreamMode `idl:"name:mode" json:"mode"`
}

func (o *SetStreamModeRequest) xxx_ToOp(ctx context.Context, op *xxx_SetStreamModeOperation) *xxx_SetStreamModeOperation {
	if op == nil {
		op = &xxx_SetStreamModeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Mode = o.Mode
	return op
}

func (o *SetStreamModeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetStreamModeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Mode = op.Mode
}
func (o *SetStreamModeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetStreamModeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetStreamModeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetStreamModeResponse structure represents the StreamMode operation response
type SetStreamModeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The StreamMode return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetStreamModeResponse) xxx_ToOp(ctx context.Context, op *xxx_SetStreamModeOperation) *xxx_SetStreamModeOperation {
	if op == nil {
		op = &xxx_SetStreamModeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetStreamModeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetStreamModeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetStreamModeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetStreamModeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetStreamModeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTraceDataProvidersOperation structure represents the TraceDataProviders operation
type xxx_GetTraceDataProvidersOperation struct {
	This      *dcom.ORPCThis                   `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat                   `idl:"name:That" json:"that"`
	Providers *pla.TraceDataProviderCollection `idl:"name:providers" json:"providers"`
	Return    int32                            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTraceDataProvidersOperation) OpNum() int { return 71 }

func (o *xxx_GetTraceDataProvidersOperation) OpName() string {
	return "/ITraceDataCollector/v0/TraceDataProviders"
}

func (o *xxx_GetTraceDataProvidersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTraceDataProvidersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTraceDataProvidersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTraceDataProvidersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTraceDataProvidersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// providers {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=ITraceDataProviderCollection}(interface))
	{
		if o.Providers != nil {
			_ptr_providers := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Providers != nil {
					if err := o.Providers.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&pla.TraceDataProviderCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Providers, _ptr_providers); err != nil {
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

func (o *xxx_GetTraceDataProvidersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// providers {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=ITraceDataProviderCollection}(interface))
	{
		_ptr_providers := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Providers == nil {
				o.Providers = &pla.TraceDataProviderCollection{}
			}
			if err := o.Providers.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_providers := func(ptr interface{}) { o.Providers = *ptr.(**pla.TraceDataProviderCollection) }
		if err := w.ReadPointer(&o.Providers, _s_providers, _ptr_providers); err != nil {
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

// GetTraceDataProvidersRequest structure represents the TraceDataProviders operation request
type GetTraceDataProvidersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTraceDataProvidersRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTraceDataProvidersOperation) *xxx_GetTraceDataProvidersOperation {
	if op == nil {
		op = &xxx_GetTraceDataProvidersOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetTraceDataProvidersRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTraceDataProvidersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTraceDataProvidersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTraceDataProvidersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTraceDataProvidersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTraceDataProvidersResponse structure represents the TraceDataProviders operation response
type GetTraceDataProvidersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// providers: Receives a pointer to the trace data provider collection object.
	Providers *pla.TraceDataProviderCollection `idl:"name:providers" json:"providers"`
	// Return: The TraceDataProviders return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTraceDataProvidersResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTraceDataProvidersOperation) *xxx_GetTraceDataProvidersOperation {
	if op == nil {
		op = &xxx_GetTraceDataProvidersOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Providers = o.Providers
	op.Return = o.Return
	return op
}

func (o *GetTraceDataProvidersResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTraceDataProvidersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Providers = op.Providers
	o.Return = op.Return
}
func (o *GetTraceDataProvidersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTraceDataProvidersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTraceDataProvidersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
