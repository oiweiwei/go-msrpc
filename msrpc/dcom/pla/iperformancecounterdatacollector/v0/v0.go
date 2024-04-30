package iperformancecounterdatacollector

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
	_ = oaut.GoPackage
	_ = pla.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/pla"
)

var (
	// IPerformanceCounterDataCollector interface identifier 03837506-098b-11d8-9414-505054503030
	PerformanceCounterDataCollectorIID = &dcom.IID{Data1: 0x03837506, Data2: 0x098b, Data3: 0x11d8, Data4: []byte{0x94, 0x14, 0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	// Syntax UUID
	PerformanceCounterDataCollectorSyntaxUUID = &uuid.UUID{TimeLow: 0x3837506, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	// Syntax ID
	PerformanceCounterDataCollectorSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: PerformanceCounterDataCollectorSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IPerformanceCounterDataCollector interface.
type PerformanceCounterDataCollectorClient interface {

	// IDataCollector retrieval method.
	DataCollector() idatacollector.DataCollectorClient

	// DataSourceName operation.
	GetDataSourceName(context.Context, *GetDataSourceNameRequest, ...dcerpc.CallOption) (*GetDataSourceNameResponse, error)

	// DataSourceName operation.
	SetDataSourceName(context.Context, *SetDataSourceNameRequest, ...dcerpc.CallOption) (*SetDataSourceNameResponse, error)

	// PerformanceCounters operation.
	GetPerformanceCounters(context.Context, *GetPerformanceCountersRequest, ...dcerpc.CallOption) (*GetPerformanceCountersResponse, error)

	// PerformanceCounters operation.
	SetPerformanceCounters(context.Context, *SetPerformanceCountersRequest, ...dcerpc.CallOption) (*SetPerformanceCountersResponse, error)

	// LogFileFormat operation.
	GetLogFileFormat(context.Context, *GetLogFileFormatRequest, ...dcerpc.CallOption) (*GetLogFileFormatResponse, error)

	// LogFileFormat operation.
	SetLogFileFormat(context.Context, *SetLogFileFormatRequest, ...dcerpc.CallOption) (*SetLogFileFormatResponse, error)

	// SampleInterval operation.
	GetSampleInterval(context.Context, *GetSampleIntervalRequest, ...dcerpc.CallOption) (*GetSampleIntervalResponse, error)

	// SampleInterval operation.
	SetSampleInterval(context.Context, *SetSampleIntervalRequest, ...dcerpc.CallOption) (*SetSampleIntervalResponse, error)

	// SegmentMaxRecords operation.
	GetSegmentMaxRecords(context.Context, *GetSegmentMaxRecordsRequest, ...dcerpc.CallOption) (*GetSegmentMaxRecordsResponse, error)

	// SegmentMaxRecords operation.
	SetSegmentMaxRecords(context.Context, *SetSegmentMaxRecordsRequest, ...dcerpc.CallOption) (*SetSegmentMaxRecordsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) PerformanceCounterDataCollectorClient
}

type xxx_DefaultPerformanceCounterDataCollectorClient struct {
	idatacollector.DataCollectorClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultPerformanceCounterDataCollectorClient) DataCollector() idatacollector.DataCollectorClient {
	return o.DataCollectorClient
}

func (o *xxx_DefaultPerformanceCounterDataCollectorClient) GetDataSourceName(ctx context.Context, in *GetDataSourceNameRequest, opts ...dcerpc.CallOption) (*GetDataSourceNameResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetDataSourceNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerformanceCounterDataCollectorClient) SetDataSourceName(ctx context.Context, in *SetDataSourceNameRequest, opts ...dcerpc.CallOption) (*SetDataSourceNameResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &SetDataSourceNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerformanceCounterDataCollectorClient) GetPerformanceCounters(ctx context.Context, in *GetPerformanceCountersRequest, opts ...dcerpc.CallOption) (*GetPerformanceCountersResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetPerformanceCountersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerformanceCounterDataCollectorClient) SetPerformanceCounters(ctx context.Context, in *SetPerformanceCountersRequest, opts ...dcerpc.CallOption) (*SetPerformanceCountersResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &SetPerformanceCountersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerformanceCounterDataCollectorClient) GetLogFileFormat(ctx context.Context, in *GetLogFileFormatRequest, opts ...dcerpc.CallOption) (*GetLogFileFormatResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetLogFileFormatResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerformanceCounterDataCollectorClient) SetLogFileFormat(ctx context.Context, in *SetLogFileFormatRequest, opts ...dcerpc.CallOption) (*SetLogFileFormatResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &SetLogFileFormatResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerformanceCounterDataCollectorClient) GetSampleInterval(ctx context.Context, in *GetSampleIntervalRequest, opts ...dcerpc.CallOption) (*GetSampleIntervalResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetSampleIntervalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerformanceCounterDataCollectorClient) SetSampleInterval(ctx context.Context, in *SetSampleIntervalRequest, opts ...dcerpc.CallOption) (*SetSampleIntervalResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &SetSampleIntervalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerformanceCounterDataCollectorClient) GetSegmentMaxRecords(ctx context.Context, in *GetSegmentMaxRecordsRequest, opts ...dcerpc.CallOption) (*GetSegmentMaxRecordsResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetSegmentMaxRecordsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerformanceCounterDataCollectorClient) SetSegmentMaxRecords(ctx context.Context, in *SetSegmentMaxRecordsRequest, opts ...dcerpc.CallOption) (*SetSegmentMaxRecordsResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &SetSegmentMaxRecordsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerformanceCounterDataCollectorClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultPerformanceCounterDataCollectorClient) IPID(ctx context.Context, ipid *dcom.IPID) PerformanceCounterDataCollectorClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultPerformanceCounterDataCollectorClient{
		DataCollectorClient: o.DataCollectorClient.IPID(ctx, ipid),
		cc:                  o.cc,
		ipid:                ipid,
	}
}
func NewPerformanceCounterDataCollectorClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (PerformanceCounterDataCollectorClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(PerformanceCounterDataCollectorSyntaxV0_0))...)
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
	return &xxx_DefaultPerformanceCounterDataCollectorClient{
		DataCollectorClient: base,
		cc:                  cc,
		ipid:                ipid,
	}, nil
}

// xxx_GetDataSourceNameOperation structure represents the DataSourceName operation
type xxx_GetDataSourceNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	DSN    *oaut.String   `idl:"name:dsn" json:"dsn"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDataSourceNameOperation) OpNum() int { return 32 }

func (o *xxx_GetDataSourceNameOperation) OpName() string {
	return "/IPerformanceCounterDataCollector/v0/DataSourceName"
}

func (o *xxx_GetDataSourceNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataSourceNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDataSourceNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDataSourceNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataSourceNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// dsn {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DSN != nil {
			_ptr_dsn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DSN != nil {
					if err := o.DSN.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DSN, _ptr_dsn); err != nil {
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

func (o *xxx_GetDataSourceNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// dsn {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_dsn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DSN == nil {
				o.DSN = &oaut.String{}
			}
			if err := o.DSN.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_dsn := func(ptr interface{}) { o.DSN = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DSN, _s_dsn, _ptr_dsn); err != nil {
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

// GetDataSourceNameRequest structure represents the DataSourceName operation request
type GetDataSourceNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDataSourceNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetDataSourceNameOperation {
	if o == nil {
		return &xxx_GetDataSourceNameOperation{}
	}
	return &xxx_GetDataSourceNameOperation{
		This: o.This,
	}
}

func (o *GetDataSourceNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDataSourceNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDataSourceNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDataSourceNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDataSourceNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDataSourceNameResponse structure represents the DataSourceName operation response
type GetDataSourceNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	DSN  *oaut.String   `idl:"name:dsn" json:"dsn"`
	// Return: The DataSourceName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDataSourceNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetDataSourceNameOperation {
	if o == nil {
		return &xxx_GetDataSourceNameOperation{}
	}
	return &xxx_GetDataSourceNameOperation{
		That:   o.That,
		DSN:    o.DSN,
		Return: o.Return,
	}
}

func (o *GetDataSourceNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDataSourceNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DSN = op.DSN
	o.Return = op.Return
}
func (o *GetDataSourceNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDataSourceNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDataSourceNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDataSourceNameOperation structure represents the DataSourceName operation
type xxx_SetDataSourceNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	DSN    *oaut.String   `idl:"name:dsn" json:"dsn"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDataSourceNameOperation) OpNum() int { return 33 }

func (o *xxx_SetDataSourceNameOperation) OpName() string {
	return "/IPerformanceCounterDataCollector/v0/DataSourceName"
}

func (o *xxx_SetDataSourceNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDataSourceNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dsn {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DSN != nil {
			_ptr_dsn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DSN != nil {
					if err := o.DSN.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DSN, _ptr_dsn); err != nil {
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

func (o *xxx_SetDataSourceNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dsn {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_dsn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DSN == nil {
				o.DSN = &oaut.String{}
			}
			if err := o.DSN.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_dsn := func(ptr interface{}) { o.DSN = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DSN, _s_dsn, _ptr_dsn); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDataSourceNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDataSourceNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetDataSourceNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetDataSourceNameRequest structure represents the DataSourceName operation request
type SetDataSourceNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	DSN  *oaut.String   `idl:"name:dsn" json:"dsn"`
}

func (o *SetDataSourceNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetDataSourceNameOperation {
	if o == nil {
		return &xxx_SetDataSourceNameOperation{}
	}
	return &xxx_SetDataSourceNameOperation{
		This: o.This,
		DSN:  o.DSN,
	}
}

func (o *SetDataSourceNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDataSourceNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DSN = op.DSN
}
func (o *SetDataSourceNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetDataSourceNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDataSourceNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDataSourceNameResponse structure represents the DataSourceName operation response
type SetDataSourceNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DataSourceName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDataSourceNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetDataSourceNameOperation {
	if o == nil {
		return &xxx_SetDataSourceNameOperation{}
	}
	return &xxx_SetDataSourceNameOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetDataSourceNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDataSourceNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDataSourceNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetDataSourceNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDataSourceNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPerformanceCountersOperation structure represents the PerformanceCounters operation
type xxx_GetPerformanceCountersOperation struct {
	This     *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Counters *oaut.SafeArray `idl:"name:counters" json:"counters"`
	Return   int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPerformanceCountersOperation) OpNum() int { return 34 }

func (o *xxx_GetPerformanceCountersOperation) OpName() string {
	return "/IPerformanceCounterDataCollector/v0/PerformanceCounters"
}

func (o *xxx_GetPerformanceCountersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPerformanceCountersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPerformanceCountersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPerformanceCountersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPerformanceCountersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// counters {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Counters != nil {
			_ptr_counters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Counters != nil {
					if err := o.Counters.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Counters, _ptr_counters); err != nil {
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

func (o *xxx_GetPerformanceCountersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// counters {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_counters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Counters == nil {
				o.Counters = &oaut.SafeArray{}
			}
			if err := o.Counters.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_counters := func(ptr interface{}) { o.Counters = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Counters, _s_counters, _ptr_counters); err != nil {
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

// GetPerformanceCountersRequest structure represents the PerformanceCounters operation request
type GetPerformanceCountersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPerformanceCountersRequest) xxx_ToOp(ctx context.Context) *xxx_GetPerformanceCountersOperation {
	if o == nil {
		return &xxx_GetPerformanceCountersOperation{}
	}
	return &xxx_GetPerformanceCountersOperation{
		This: o.This,
	}
}

func (o *GetPerformanceCountersRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPerformanceCountersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPerformanceCountersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetPerformanceCountersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPerformanceCountersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPerformanceCountersResponse structure represents the PerformanceCounters operation response
type GetPerformanceCountersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Counters *oaut.SafeArray `idl:"name:counters" json:"counters"`
	// Return: The PerformanceCounters return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPerformanceCountersResponse) xxx_ToOp(ctx context.Context) *xxx_GetPerformanceCountersOperation {
	if o == nil {
		return &xxx_GetPerformanceCountersOperation{}
	}
	return &xxx_GetPerformanceCountersOperation{
		That:     o.That,
		Counters: o.Counters,
		Return:   o.Return,
	}
}

func (o *GetPerformanceCountersResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPerformanceCountersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Counters = op.Counters
	o.Return = op.Return
}
func (o *GetPerformanceCountersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetPerformanceCountersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPerformanceCountersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPerformanceCountersOperation structure represents the PerformanceCounters operation
type xxx_SetPerformanceCountersOperation struct {
	This     *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Counters *oaut.SafeArray `idl:"name:counters" json:"counters"`
	Return   int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPerformanceCountersOperation) OpNum() int { return 35 }

func (o *xxx_SetPerformanceCountersOperation) OpName() string {
	return "/IPerformanceCounterDataCollector/v0/PerformanceCounters"
}

func (o *xxx_SetPerformanceCountersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPerformanceCountersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// counters {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Counters != nil {
			_ptr_counters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Counters != nil {
					if err := o.Counters.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Counters, _ptr_counters); err != nil {
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

func (o *xxx_SetPerformanceCountersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// counters {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_counters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Counters == nil {
				o.Counters = &oaut.SafeArray{}
			}
			if err := o.Counters.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_counters := func(ptr interface{}) { o.Counters = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Counters, _s_counters, _ptr_counters); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPerformanceCountersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPerformanceCountersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetPerformanceCountersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetPerformanceCountersRequest structure represents the PerformanceCounters operation request
type SetPerformanceCountersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis  `idl:"name:This" json:"this"`
	Counters *oaut.SafeArray `idl:"name:counters" json:"counters"`
}

func (o *SetPerformanceCountersRequest) xxx_ToOp(ctx context.Context) *xxx_SetPerformanceCountersOperation {
	if o == nil {
		return &xxx_SetPerformanceCountersOperation{}
	}
	return &xxx_SetPerformanceCountersOperation{
		This:     o.This,
		Counters: o.Counters,
	}
}

func (o *SetPerformanceCountersRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPerformanceCountersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Counters = op.Counters
}
func (o *SetPerformanceCountersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetPerformanceCountersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPerformanceCountersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPerformanceCountersResponse structure represents the PerformanceCounters operation response
type SetPerformanceCountersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PerformanceCounters return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPerformanceCountersResponse) xxx_ToOp(ctx context.Context) *xxx_SetPerformanceCountersOperation {
	if o == nil {
		return &xxx_SetPerformanceCountersOperation{}
	}
	return &xxx_SetPerformanceCountersOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetPerformanceCountersResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPerformanceCountersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetPerformanceCountersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetPerformanceCountersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPerformanceCountersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLogFileFormatOperation structure represents the LogFileFormat operation
type xxx_GetLogFileFormatOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Format pla.FileFormat `idl:"name:format" json:"format"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLogFileFormatOperation) OpNum() int { return 36 }

func (o *xxx_GetLogFileFormatOperation) OpName() string {
	return "/IPerformanceCounterDataCollector/v0/LogFileFormat"
}

func (o *xxx_GetLogFileFormatOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogFileFormatOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLogFileFormatOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLogFileFormatOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogFileFormatOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// format {out, retval} (1:{pointer=ref}*(1))(2:{alias=FileFormat}(enum))
	{
		if err := w.WriteData(uint16(o.Format)); err != nil {
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

func (o *xxx_GetLogFileFormatOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// format {out, retval} (1:{pointer=ref}*(1))(2:{alias=FileFormat}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Format)); err != nil {
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

// GetLogFileFormatRequest structure represents the LogFileFormat operation request
type GetLogFileFormatRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLogFileFormatRequest) xxx_ToOp(ctx context.Context) *xxx_GetLogFileFormatOperation {
	if o == nil {
		return &xxx_GetLogFileFormatOperation{}
	}
	return &xxx_GetLogFileFormatOperation{
		This: o.This,
	}
}

func (o *GetLogFileFormatRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLogFileFormatOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLogFileFormatRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetLogFileFormatRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLogFileFormatOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLogFileFormatResponse structure represents the LogFileFormat operation response
type GetLogFileFormatResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Format pla.FileFormat `idl:"name:format" json:"format"`
	// Return: The LogFileFormat return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLogFileFormatResponse) xxx_ToOp(ctx context.Context) *xxx_GetLogFileFormatOperation {
	if o == nil {
		return &xxx_GetLogFileFormatOperation{}
	}
	return &xxx_GetLogFileFormatOperation{
		That:   o.That,
		Format: o.Format,
		Return: o.Return,
	}
}

func (o *GetLogFileFormatResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLogFileFormatOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Format = op.Format
	o.Return = op.Return
}
func (o *GetLogFileFormatResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetLogFileFormatResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLogFileFormatOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetLogFileFormatOperation structure represents the LogFileFormat operation
type xxx_SetLogFileFormatOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Format pla.FileFormat `idl:"name:format" json:"format"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetLogFileFormatOperation) OpNum() int { return 37 }

func (o *xxx_SetLogFileFormatOperation) OpName() string {
	return "/IPerformanceCounterDataCollector/v0/LogFileFormat"
}

func (o *xxx_SetLogFileFormatOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogFileFormatOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// format {in} (1:{alias=FileFormat}(enum))
	{
		if err := w.WriteData(uint16(o.Format)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogFileFormatOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// format {in} (1:{alias=FileFormat}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Format)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogFileFormatOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogFileFormatOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetLogFileFormatOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetLogFileFormatRequest structure represents the LogFileFormat operation request
type SetLogFileFormatRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	Format pla.FileFormat `idl:"name:format" json:"format"`
}

func (o *SetLogFileFormatRequest) xxx_ToOp(ctx context.Context) *xxx_SetLogFileFormatOperation {
	if o == nil {
		return &xxx_SetLogFileFormatOperation{}
	}
	return &xxx_SetLogFileFormatOperation{
		This:   o.This,
		Format: o.Format,
	}
}

func (o *SetLogFileFormatRequest) xxx_FromOp(ctx context.Context, op *xxx_SetLogFileFormatOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Format = op.Format
}
func (o *SetLogFileFormatRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetLogFileFormatRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLogFileFormatOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetLogFileFormatResponse structure represents the LogFileFormat operation response
type SetLogFileFormatResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The LogFileFormat return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetLogFileFormatResponse) xxx_ToOp(ctx context.Context) *xxx_SetLogFileFormatOperation {
	if o == nil {
		return &xxx_SetLogFileFormatOperation{}
	}
	return &xxx_SetLogFileFormatOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetLogFileFormatResponse) xxx_FromOp(ctx context.Context, op *xxx_SetLogFileFormatOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetLogFileFormatResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetLogFileFormatResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLogFileFormatOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSampleIntervalOperation structure represents the SampleInterval operation
type xxx_GetSampleIntervalOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Interval uint32         `idl:"name:interval" json:"interval"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSampleIntervalOperation) OpNum() int { return 38 }

func (o *xxx_GetSampleIntervalOperation) OpName() string {
	return "/IPerformanceCounterDataCollector/v0/SampleInterval"
}

func (o *xxx_GetSampleIntervalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSampleIntervalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSampleIntervalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSampleIntervalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSampleIntervalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// interval {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Interval); err != nil {
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

func (o *xxx_GetSampleIntervalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// interval {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Interval); err != nil {
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

// GetSampleIntervalRequest structure represents the SampleInterval operation request
type GetSampleIntervalRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSampleIntervalRequest) xxx_ToOp(ctx context.Context) *xxx_GetSampleIntervalOperation {
	if o == nil {
		return &xxx_GetSampleIntervalOperation{}
	}
	return &xxx_GetSampleIntervalOperation{
		This: o.This,
	}
}

func (o *GetSampleIntervalRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSampleIntervalOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSampleIntervalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSampleIntervalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSampleIntervalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSampleIntervalResponse structure represents the SampleInterval operation response
type GetSampleIntervalResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Interval uint32         `idl:"name:interval" json:"interval"`
	// Return: The SampleInterval return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSampleIntervalResponse) xxx_ToOp(ctx context.Context) *xxx_GetSampleIntervalOperation {
	if o == nil {
		return &xxx_GetSampleIntervalOperation{}
	}
	return &xxx_GetSampleIntervalOperation{
		That:     o.That,
		Interval: o.Interval,
		Return:   o.Return,
	}
}

func (o *GetSampleIntervalResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSampleIntervalOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Interval = op.Interval
	o.Return = op.Return
}
func (o *GetSampleIntervalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSampleIntervalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSampleIntervalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSampleIntervalOperation structure represents the SampleInterval operation
type xxx_SetSampleIntervalOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Interval uint32         `idl:"name:interval" json:"interval"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSampleIntervalOperation) OpNum() int { return 39 }

func (o *xxx_SetSampleIntervalOperation) OpName() string {
	return "/IPerformanceCounterDataCollector/v0/SampleInterval"
}

func (o *xxx_SetSampleIntervalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSampleIntervalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// interval {in} (1:(uint32))
	{
		if err := w.WriteData(o.Interval); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSampleIntervalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// interval {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Interval); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSampleIntervalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSampleIntervalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSampleIntervalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSampleIntervalRequest structure represents the SampleInterval operation request
type SetSampleIntervalRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	Interval uint32         `idl:"name:interval" json:"interval"`
}

func (o *SetSampleIntervalRequest) xxx_ToOp(ctx context.Context) *xxx_SetSampleIntervalOperation {
	if o == nil {
		return &xxx_SetSampleIntervalOperation{}
	}
	return &xxx_SetSampleIntervalOperation{
		This:     o.This,
		Interval: o.Interval,
	}
}

func (o *SetSampleIntervalRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSampleIntervalOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Interval = op.Interval
}
func (o *SetSampleIntervalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSampleIntervalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSampleIntervalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSampleIntervalResponse structure represents the SampleInterval operation response
type SetSampleIntervalResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SampleInterval return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSampleIntervalResponse) xxx_ToOp(ctx context.Context) *xxx_SetSampleIntervalOperation {
	if o == nil {
		return &xxx_SetSampleIntervalOperation{}
	}
	return &xxx_SetSampleIntervalOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSampleIntervalResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSampleIntervalOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSampleIntervalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSampleIntervalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSampleIntervalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSegmentMaxRecordsOperation structure represents the SegmentMaxRecords operation
type xxx_GetSegmentMaxRecordsOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Records uint32         `idl:"name:records" json:"records"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSegmentMaxRecordsOperation) OpNum() int { return 40 }

func (o *xxx_GetSegmentMaxRecordsOperation) OpName() string {
	return "/IPerformanceCounterDataCollector/v0/SegmentMaxRecords"
}

func (o *xxx_GetSegmentMaxRecordsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSegmentMaxRecordsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSegmentMaxRecordsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSegmentMaxRecordsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSegmentMaxRecordsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// records {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Records); err != nil {
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

func (o *xxx_GetSegmentMaxRecordsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// records {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Records); err != nil {
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

// GetSegmentMaxRecordsRequest structure represents the SegmentMaxRecords operation request
type GetSegmentMaxRecordsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSegmentMaxRecordsRequest) xxx_ToOp(ctx context.Context) *xxx_GetSegmentMaxRecordsOperation {
	if o == nil {
		return &xxx_GetSegmentMaxRecordsOperation{}
	}
	return &xxx_GetSegmentMaxRecordsOperation{
		This: o.This,
	}
}

func (o *GetSegmentMaxRecordsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSegmentMaxRecordsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSegmentMaxRecordsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSegmentMaxRecordsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSegmentMaxRecordsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSegmentMaxRecordsResponse structure represents the SegmentMaxRecords operation response
type GetSegmentMaxRecordsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Records uint32         `idl:"name:records" json:"records"`
	// Return: The SegmentMaxRecords return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSegmentMaxRecordsResponse) xxx_ToOp(ctx context.Context) *xxx_GetSegmentMaxRecordsOperation {
	if o == nil {
		return &xxx_GetSegmentMaxRecordsOperation{}
	}
	return &xxx_GetSegmentMaxRecordsOperation{
		That:    o.That,
		Records: o.Records,
		Return:  o.Return,
	}
}

func (o *GetSegmentMaxRecordsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSegmentMaxRecordsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Records = op.Records
	o.Return = op.Return
}
func (o *GetSegmentMaxRecordsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSegmentMaxRecordsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSegmentMaxRecordsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSegmentMaxRecordsOperation structure represents the SegmentMaxRecords operation
type xxx_SetSegmentMaxRecordsOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Records uint32         `idl:"name:records" json:"records"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSegmentMaxRecordsOperation) OpNum() int { return 41 }

func (o *xxx_SetSegmentMaxRecordsOperation) OpName() string {
	return "/IPerformanceCounterDataCollector/v0/SegmentMaxRecords"
}

func (o *xxx_SetSegmentMaxRecordsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentMaxRecordsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// records {in} (1:(uint32))
	{
		if err := w.WriteData(o.Records); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentMaxRecordsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// records {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Records); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentMaxRecordsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentMaxRecordsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSegmentMaxRecordsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSegmentMaxRecordsRequest structure represents the SegmentMaxRecords operation request
type SetSegmentMaxRecordsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Records uint32         `idl:"name:records" json:"records"`
}

func (o *SetSegmentMaxRecordsRequest) xxx_ToOp(ctx context.Context) *xxx_SetSegmentMaxRecordsOperation {
	if o == nil {
		return &xxx_SetSegmentMaxRecordsOperation{}
	}
	return &xxx_SetSegmentMaxRecordsOperation{
		This:    o.This,
		Records: o.Records,
	}
}

func (o *SetSegmentMaxRecordsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSegmentMaxRecordsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Records = op.Records
}
func (o *SetSegmentMaxRecordsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSegmentMaxRecordsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSegmentMaxRecordsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSegmentMaxRecordsResponse structure represents the SegmentMaxRecords operation response
type SetSegmentMaxRecordsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SegmentMaxRecords return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSegmentMaxRecordsResponse) xxx_ToOp(ctx context.Context) *xxx_SetSegmentMaxRecordsOperation {
	if o == nil {
		return &xxx_SetSegmentMaxRecordsOperation{}
	}
	return &xxx_SetSegmentMaxRecordsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSegmentMaxRecordsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSegmentMaxRecordsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSegmentMaxRecordsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSegmentMaxRecordsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSegmentMaxRecordsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
