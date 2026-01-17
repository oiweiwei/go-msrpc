package itracedatacollector

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
		op := &xxx_GetBufferSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBufferSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBufferSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // BufferSize
		op := &xxx_SetBufferSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetBufferSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetBufferSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // BuffersLost
		op := &xxx_GetBuffersLostOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBuffersLostRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBuffersLost(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // Opnum35NotUsedOnWire
		// Opnum35NotUsedOnWire
		return nil, nil
	case 36: // BuffersWritten
		op := &xxx_GetBuffersWrittenOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBuffersWrittenRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBuffersWritten(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // Opnum37NotUsedOnWire
		// Opnum37NotUsedOnWire
		return nil, nil
	case 38: // ClockType
		op := &xxx_GetClockTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClockTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClockType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // ClockType
		op := &xxx_SetClockTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetClockTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetClockType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // EventsLost
		op := &xxx_GetEventsLostOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEventsLostRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEventsLost(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // Opnum41NotUsedOnWire
		// Opnum41NotUsedOnWire
		return nil, nil
	case 42: // ExtendedModes
		op := &xxx_GetExtendedModesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetExtendedModesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetExtendedModes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // ExtendedModes
		op := &xxx_SetExtendedModesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetExtendedModesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetExtendedModes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // FlushTimer
		op := &xxx_GetFlushTimerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFlushTimerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFlushTimer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // FlushTimer
		op := &xxx_SetFlushTimerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFlushTimerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFlushTimer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // FreeBuffers
		op := &xxx_GetFreeBuffersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFreeBuffersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFreeBuffers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // Opnum47NotUsedOnWire
		// Opnum47NotUsedOnWire
		return nil, nil
	case 48: // Guid
		op := &xxx_GetGUIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetGUIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetGUID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // Guid
		op := &xxx_SetGUIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetGUIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetGUID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 50: // IsKernelTrace
		op := &xxx_GetIsKernelTraceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsKernelTraceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsKernelTrace(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 51: // MaximumBuffers
		op := &xxx_GetMaximumBuffersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMaximumBuffersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMaximumBuffers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 52: // MaximumBuffers
		op := &xxx_SetMaximumBuffersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMaximumBuffersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMaximumBuffers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 53: // MinimumBuffers
		op := &xxx_GetMinimumBuffersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMinimumBuffersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMinimumBuffers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 54: // MinimumBuffers
		op := &xxx_SetMinimumBuffersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMinimumBuffersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMinimumBuffers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 55: // NumberOfBuffers
		op := &xxx_GetNumberOfBuffersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNumberOfBuffersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNumberOfBuffers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 56: // NumberOfBuffers
		op := &xxx_SetNumberOfBuffersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNumberOfBuffersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNumberOfBuffers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 57: // PreallocateFile
		op := &xxx_GetPreallocateFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPreallocateFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPreallocateFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 58: // PreallocateFile
		op := &xxx_SetPreallocateFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPreallocateFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPreallocateFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 59: // ProcessMode
		op := &xxx_GetProcessModeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetProcessModeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProcessMode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 60: // ProcessMode
		op := &xxx_SetProcessModeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetProcessModeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetProcessMode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 61: // RealTimeBuffersLost
		op := &xxx_GetRealTimeBuffersLostOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRealTimeBuffersLostRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRealTimeBuffersLost(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 62: // Opnum62NotUsedOnWire
		// Opnum62NotUsedOnWire
		return nil, nil
	case 63: // SessionId
		op := &xxx_GetSessionIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSessionIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSessionID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 64: // Opnum64NotUsedOnWire
		// Opnum64NotUsedOnWire
		return nil, nil
	case 65: // SessionName
		op := &xxx_GetSessionNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSessionNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSessionName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 66: // SessionName
		op := &xxx_SetSessionNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSessionNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSessionName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 67: // SessionThreadId
		op := &xxx_GetSessionThreadIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSessionThreadIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSessionThreadID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 68: // Opnum68NotUsedOnWire
		// Opnum68NotUsedOnWire
		return nil, nil
	case 69: // StreamMode
		op := &xxx_GetStreamModeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetStreamModeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetStreamMode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 70: // StreamMode
		op := &xxx_SetStreamModeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetStreamModeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetStreamMode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 71: // TraceDataProviders
		op := &xxx_GetTraceDataProvidersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTraceDataProvidersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTraceDataProviders(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ITraceDataCollector
type UnimplementedTraceDataCollectorServer struct {
	idatacollector.UnimplementedDataCollectorServer
}

func (UnimplementedTraceDataCollectorServer) GetBufferSize(context.Context, *GetBufferSizeRequest) (*GetBufferSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) SetBufferSize(context.Context, *SetBufferSizeRequest) (*SetBufferSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetBuffersLost(context.Context, *GetBuffersLostRequest) (*GetBuffersLostResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetBuffersWritten(context.Context, *GetBuffersWrittenRequest) (*GetBuffersWrittenResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetClockType(context.Context, *GetClockTypeRequest) (*GetClockTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) SetClockType(context.Context, *SetClockTypeRequest) (*SetClockTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetEventsLost(context.Context, *GetEventsLostRequest) (*GetEventsLostResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetExtendedModes(context.Context, *GetExtendedModesRequest) (*GetExtendedModesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) SetExtendedModes(context.Context, *SetExtendedModesRequest) (*SetExtendedModesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetFlushTimer(context.Context, *GetFlushTimerRequest) (*GetFlushTimerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) SetFlushTimer(context.Context, *SetFlushTimerRequest) (*SetFlushTimerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetFreeBuffers(context.Context, *GetFreeBuffersRequest) (*GetFreeBuffersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetGUID(context.Context, *GetGUIDRequest) (*GetGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) SetGUID(context.Context, *SetGUIDRequest) (*SetGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetIsKernelTrace(context.Context, *GetIsKernelTraceRequest) (*GetIsKernelTraceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetMaximumBuffers(context.Context, *GetMaximumBuffersRequest) (*GetMaximumBuffersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) SetMaximumBuffers(context.Context, *SetMaximumBuffersRequest) (*SetMaximumBuffersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetMinimumBuffers(context.Context, *GetMinimumBuffersRequest) (*GetMinimumBuffersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) SetMinimumBuffers(context.Context, *SetMinimumBuffersRequest) (*SetMinimumBuffersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetNumberOfBuffers(context.Context, *GetNumberOfBuffersRequest) (*GetNumberOfBuffersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) SetNumberOfBuffers(context.Context, *SetNumberOfBuffersRequest) (*SetNumberOfBuffersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetPreallocateFile(context.Context, *GetPreallocateFileRequest) (*GetPreallocateFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) SetPreallocateFile(context.Context, *SetPreallocateFileRequest) (*SetPreallocateFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetProcessMode(context.Context, *GetProcessModeRequest) (*GetProcessModeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) SetProcessMode(context.Context, *SetProcessModeRequest) (*SetProcessModeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetRealTimeBuffersLost(context.Context, *GetRealTimeBuffersLostRequest) (*GetRealTimeBuffersLostResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetSessionID(context.Context, *GetSessionIDRequest) (*GetSessionIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetSessionName(context.Context, *GetSessionNameRequest) (*GetSessionNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) SetSessionName(context.Context, *SetSessionNameRequest) (*SetSessionNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetSessionThreadID(context.Context, *GetSessionThreadIDRequest) (*GetSessionThreadIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetStreamMode(context.Context, *GetStreamModeRequest) (*GetStreamModeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) SetStreamMode(context.Context, *SetStreamModeRequest) (*SetStreamModeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataCollectorServer) GetTraceDataProviders(context.Context, *GetTraceDataProvidersRequest) (*GetTraceDataProvidersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TraceDataCollectorServer = (*UnimplementedTraceDataCollectorServer)(nil)
