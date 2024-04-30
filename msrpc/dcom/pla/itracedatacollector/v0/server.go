package itracedatacollector

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

// ITraceDataCollector server interface.
type TraceDataCollectorServer interface {

	// IDataCollector base class.
	idatacollector.DataCollectorServer

	// BufferSize operation.
	GetBufferSize(context.Context, *GetBufferSizeRequest) (*GetBufferSizeResponse, error)

	// BufferSize operation.
	SetBufferSize(context.Context, *SetBufferSizeRequest) (*SetBufferSizeResponse, error)

	// The BuffersLost (Get) method retrieves the BufferLost property, as specified in the
	// property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetBuffersLost(context.Context, *GetBuffersLostRequest) (*GetBuffersLostResponse, error)

	// Opnum35NotUsedOnWire operation.
	// Opnum35NotUsedOnWire

	// The BuffersWritten (Get) method retrieves the BuffersWritten property, as specified
	// in the property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetBuffersWritten(context.Context, *GetBuffersWrittenRequest) (*GetBuffersWrittenResponse, error)

	// Opnum37NotUsedOnWire operation.
	// Opnum37NotUsedOnWire

	// The ClockType enumeration defines the clock resolution to use when tracing events.
	GetClockType(context.Context, *GetClockTypeRequest) (*GetClockTypeResponse, error)

	// The ClockType enumeration defines the clock resolution to use when tracing events.
	SetClockType(context.Context, *SetClockTypeRequest) (*SetClockTypeResponse, error)

	// The EventsLost (Get) method retrieves the EventsLost property, as specified in the
	// property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetEventsLost(context.Context, *GetEventsLostRequest) (*GetEventsLostResponse, error)

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
	GetExtendedModes(context.Context, *GetExtendedModesRequest) (*GetExtendedModesResponse, error)

	// The ExtendedModes (Put) method sets the ExtendedModes property, as specified in the
	// property table in section 3.2.4.9.
	//
	// The ExtendedModes (Get) method retrieves the ExtendedModes property, as specified
	// in the property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	SetExtendedModes(context.Context, *SetExtendedModesRequest) (*SetExtendedModesResponse, error)

	// FlushTimer operation.
	GetFlushTimer(context.Context, *GetFlushTimerRequest) (*GetFlushTimerResponse, error)

	// FlushTimer operation.
	SetFlushTimer(context.Context, *SetFlushTimerRequest) (*SetFlushTimerResponse, error)

	// The FreeBuffers (Get) method retrieves the FreeBuffers property, as specified in
	// the property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetFreeBuffers(context.Context, *GetFreeBuffersRequest) (*GetFreeBuffersResponse, error)

	// Opnum47NotUsedOnWire operation.
	// Opnum47NotUsedOnWire

	// Guid operation.
	GetGUID(context.Context, *GetGUIDRequest) (*GetGUIDResponse, error)

	// Guid operation.
	SetGUID(context.Context, *SetGUIDRequest) (*SetGUIDResponse, error)

	// The IsKernelTrace (Get) method retrieves the IsKernelTrace property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetIsKernelTrace(context.Context, *GetIsKernelTraceRequest) (*GetIsKernelTraceResponse, error)

	// MaximumBuffers operation.
	GetMaximumBuffers(context.Context, *GetMaximumBuffersRequest) (*GetMaximumBuffersResponse, error)

	// MaximumBuffers operation.
	SetMaximumBuffers(context.Context, *SetMaximumBuffersRequest) (*SetMaximumBuffersResponse, error)

	// MinimumBuffers operation.
	GetMinimumBuffers(context.Context, *GetMinimumBuffersRequest) (*GetMinimumBuffersResponse, error)

	// MinimumBuffers operation.
	SetMinimumBuffers(context.Context, *SetMinimumBuffersRequest) (*SetMinimumBuffersResponse, error)

	// NumberOfBuffers operation.
	GetNumberOfBuffers(context.Context, *GetNumberOfBuffersRequest) (*GetNumberOfBuffersResponse, error)

	// NumberOfBuffers operation.
	SetNumberOfBuffers(context.Context, *SetNumberOfBuffersRequest) (*SetNumberOfBuffersResponse, error)

	// PreallocateFile operation.
	GetPreallocateFile(context.Context, *GetPreallocateFileRequest) (*GetPreallocateFileResponse, error)

	// PreallocateFile operation.
	SetPreallocateFile(context.Context, *SetPreallocateFileRequest) (*SetPreallocateFileResponse, error)

	// ProcessMode operation.
	GetProcessMode(context.Context, *GetProcessModeRequest) (*GetProcessModeResponse, error)

	// ProcessMode operation.
	SetProcessMode(context.Context, *SetProcessModeRequest) (*SetProcessModeResponse, error)

	// The RealTimeBuffersLost (Get) method retrieves the RealTimeBuffersLost property,
	// as specified in the property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetRealTimeBuffersLost(context.Context, *GetRealTimeBuffersLostRequest) (*GetRealTimeBuffersLostResponse, error)

	// Opnum62NotUsedOnWire operation.
	// Opnum62NotUsedOnWire

	// The SessionId (Get) method retrieves the SessionId property, as specified in the
	// property table in section 3.2.4.9.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetSessionID(context.Context, *GetSessionIDRequest) (*GetSessionIDResponse, error)

	// Opnum64NotUsedOnWire operation.
	// Opnum64NotUsedOnWire

	// SessionName operation.
	GetSessionName(context.Context, *GetSessionNameRequest) (*GetSessionNameResponse, error)

	// SessionName operation.
	SetSessionName(context.Context, *SetSessionNameRequest) (*SetSessionNameResponse, error)

	// The SessionThreadId (Get) method retrieves the SessionThreadId property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetSessionThreadID(context.Context, *GetSessionThreadIDRequest) (*GetSessionThreadIDResponse, error)

	// Opnum68NotUsedOnWire operation.
	// Opnum68NotUsedOnWire

	// The StreamMode enumeration defines where the trace events are delivered.
	GetStreamMode(context.Context, *GetStreamModeRequest) (*GetStreamModeResponse, error)

	// The StreamMode enumeration defines where the trace events are delivered.
	SetStreamMode(context.Context, *SetStreamModeRequest) (*SetStreamModeResponse, error)

	// The TraceDataProviders (Get) method retrieves the TraceDataProviders property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetTraceDataProviders(context.Context, *GetTraceDataProvidersRequest) (*GetTraceDataProvidersResponse, error)
}

func RegisterTraceDataCollectorServer(conn dcerpc.Conn, o TraceDataCollectorServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTraceDataCollectorServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TraceDataCollectorSyntaxV0_0))...)
}

func NewTraceDataCollectorServerHandle(o TraceDataCollectorServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TraceDataCollectorServerHandle(ctx, o, opNum, r)
	}
}

func TraceDataCollectorServerHandle(ctx context.Context, o TraceDataCollectorServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 32 {
		// IDataCollector base method.
		return idatacollector.DataCollectorServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 32: // BufferSize
		in := &GetBufferSizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetBufferSize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // BufferSize
		in := &SetBufferSizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetBufferSize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // BuffersLost
		in := &GetBuffersLostRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetBuffersLost(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // Opnum35NotUsedOnWire
		// Opnum35NotUsedOnWire
		return nil, nil
	case 36: // BuffersWritten
		in := &GetBuffersWrittenRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetBuffersWritten(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // Opnum37NotUsedOnWire
		// Opnum37NotUsedOnWire
		return nil, nil
	case 38: // ClockType
		in := &GetClockTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClockType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // ClockType
		in := &SetClockTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetClockType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // EventsLost
		in := &GetEventsLostRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventsLost(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // Opnum41NotUsedOnWire
		// Opnum41NotUsedOnWire
		return nil, nil
	case 42: // ExtendedModes
		in := &GetExtendedModesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetExtendedModes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // ExtendedModes
		in := &SetExtendedModesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetExtendedModes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // FlushTimer
		in := &GetFlushTimerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFlushTimer(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // FlushTimer
		in := &SetFlushTimerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFlushTimer(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // FreeBuffers
		in := &GetFreeBuffersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFreeBuffers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // Opnum47NotUsedOnWire
		// Opnum47NotUsedOnWire
		return nil, nil
	case 48: // Guid
		in := &GetGUIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetGUID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // Guid
		in := &SetGUIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetGUID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 50: // IsKernelTrace
		in := &GetIsKernelTraceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIsKernelTrace(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 51: // MaximumBuffers
		in := &GetMaximumBuffersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMaximumBuffers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 52: // MaximumBuffers
		in := &SetMaximumBuffersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMaximumBuffers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 53: // MinimumBuffers
		in := &GetMinimumBuffersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMinimumBuffers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 54: // MinimumBuffers
		in := &SetMinimumBuffersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMinimumBuffers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 55: // NumberOfBuffers
		in := &GetNumberOfBuffersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNumberOfBuffers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 56: // NumberOfBuffers
		in := &SetNumberOfBuffersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetNumberOfBuffers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 57: // PreallocateFile
		in := &GetPreallocateFileRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPreallocateFile(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 58: // PreallocateFile
		in := &SetPreallocateFileRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPreallocateFile(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 59: // ProcessMode
		in := &GetProcessModeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProcessMode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 60: // ProcessMode
		in := &SetProcessModeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetProcessMode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 61: // RealTimeBuffersLost
		in := &GetRealTimeBuffersLostRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRealTimeBuffersLost(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 62: // Opnum62NotUsedOnWire
		// Opnum62NotUsedOnWire
		return nil, nil
	case 63: // SessionId
		in := &GetSessionIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSessionID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 64: // Opnum64NotUsedOnWire
		// Opnum64NotUsedOnWire
		return nil, nil
	case 65: // SessionName
		in := &GetSessionNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSessionName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 66: // SessionName
		in := &SetSessionNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSessionName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 67: // SessionThreadId
		in := &GetSessionThreadIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSessionThreadID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 68: // Opnum68NotUsedOnWire
		// Opnum68NotUsedOnWire
		return nil, nil
	case 69: // StreamMode
		in := &GetStreamModeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetStreamMode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 70: // StreamMode
		in := &SetStreamModeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetStreamMode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 71: // TraceDataProviders
		in := &GetTraceDataProvidersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTraceDataProviders(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
