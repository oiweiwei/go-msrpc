package iapitracingdatacollector

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
)

var (
	// import guard
	GoPackage = "dcom/pla"
)

var (
	// IApiTracingDataCollector interface identifier 0383751a-098b-11d8-9414-505054503030
	APITracingDataCollectorIID = &dcom.IID{Data1: 0x0383751a, Data2: 0x098b, Data3: 0x11d8, Data4: []byte{0x94, 0x14, 0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	// Syntax UUID
	APITracingDataCollectorSyntaxUUID = &uuid.UUID{TimeLow: 0x383751a, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	// Syntax ID
	APITracingDataCollectorSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: APITracingDataCollectorSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IApiTracingDataCollector interface.
type APITracingDataCollectorClient interface {

	// IDataCollector retrieval method.
	DataCollector() idatacollector.DataCollectorClient

	// LogApiNamesOnly operation.
	GetLogAPINamesOnly(context.Context, *GetLogAPINamesOnlyRequest, ...dcerpc.CallOption) (*GetLogAPINamesOnlyResponse, error)

	// LogApiNamesOnly operation.
	SetLogAPINamesOnly(context.Context, *SetLogAPINamesOnlyRequest, ...dcerpc.CallOption) (*SetLogAPINamesOnlyResponse, error)

	// LogApisRecursively operation.
	GetLogAPIsRecursively(context.Context, *GetLogAPIsRecursivelyRequest, ...dcerpc.CallOption) (*GetLogAPIsRecursivelyResponse, error)

	// LogApisRecursively operation.
	SetLogAPIsRecursively(context.Context, *SetLogAPIsRecursivelyRequest, ...dcerpc.CallOption) (*SetLogAPIsRecursivelyResponse, error)

	// ExePath operation.
	GetExePath(context.Context, *GetExePathRequest, ...dcerpc.CallOption) (*GetExePathResponse, error)

	// ExePath operation.
	SetExePath(context.Context, *SetExePathRequest, ...dcerpc.CallOption) (*SetExePathResponse, error)

	// LogFilePath operation.
	GetLogFilePath(context.Context, *GetLogFilePathRequest, ...dcerpc.CallOption) (*GetLogFilePathResponse, error)

	// LogFilePath operation.
	SetLogFilePath(context.Context, *SetLogFilePathRequest, ...dcerpc.CallOption) (*SetLogFilePathResponse, error)

	// IncludeModules operation.
	GetIncludeModules(context.Context, *GetIncludeModulesRequest, ...dcerpc.CallOption) (*GetIncludeModulesResponse, error)

	// IncludeModules operation.
	SetIncludeModules(context.Context, *SetIncludeModulesRequest, ...dcerpc.CallOption) (*SetIncludeModulesResponse, error)

	// IncludeApis operation.
	GetIncludeAPIs(context.Context, *GetIncludeAPIsRequest, ...dcerpc.CallOption) (*GetIncludeAPIsResponse, error)

	// IncludeApis operation.
	SetIncludeAPIs(context.Context, *SetIncludeAPIsRequest, ...dcerpc.CallOption) (*SetIncludeAPIsResponse, error)

	// ExcludeApis operation.
	GetExcludeAPIs(context.Context, *GetExcludeAPIsRequest, ...dcerpc.CallOption) (*GetExcludeAPIsResponse, error)

	// ExcludeApis operation.
	SetExcludeAPIs(context.Context, *SetExcludeAPIsRequest, ...dcerpc.CallOption) (*SetExcludeAPIsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) APITracingDataCollectorClient
}

type xxx_DefaultAPITracingDataCollectorClient struct {
	idatacollector.DataCollectorClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAPITracingDataCollectorClient) DataCollector() idatacollector.DataCollectorClient {
	return o.DataCollectorClient
}

func (o *xxx_DefaultAPITracingDataCollectorClient) GetLogAPINamesOnly(ctx context.Context, in *GetLogAPINamesOnlyRequest, opts ...dcerpc.CallOption) (*GetLogAPINamesOnlyResponse, error) {
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
	out := &GetLogAPINamesOnlyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAPITracingDataCollectorClient) SetLogAPINamesOnly(ctx context.Context, in *SetLogAPINamesOnlyRequest, opts ...dcerpc.CallOption) (*SetLogAPINamesOnlyResponse, error) {
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
	out := &SetLogAPINamesOnlyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAPITracingDataCollectorClient) GetLogAPIsRecursively(ctx context.Context, in *GetLogAPIsRecursivelyRequest, opts ...dcerpc.CallOption) (*GetLogAPIsRecursivelyResponse, error) {
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
	out := &GetLogAPIsRecursivelyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAPITracingDataCollectorClient) SetLogAPIsRecursively(ctx context.Context, in *SetLogAPIsRecursivelyRequest, opts ...dcerpc.CallOption) (*SetLogAPIsRecursivelyResponse, error) {
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
	out := &SetLogAPIsRecursivelyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAPITracingDataCollectorClient) GetExePath(ctx context.Context, in *GetExePathRequest, opts ...dcerpc.CallOption) (*GetExePathResponse, error) {
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
	out := &GetExePathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAPITracingDataCollectorClient) SetExePath(ctx context.Context, in *SetExePathRequest, opts ...dcerpc.CallOption) (*SetExePathResponse, error) {
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
	out := &SetExePathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAPITracingDataCollectorClient) GetLogFilePath(ctx context.Context, in *GetLogFilePathRequest, opts ...dcerpc.CallOption) (*GetLogFilePathResponse, error) {
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
	out := &GetLogFilePathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAPITracingDataCollectorClient) SetLogFilePath(ctx context.Context, in *SetLogFilePathRequest, opts ...dcerpc.CallOption) (*SetLogFilePathResponse, error) {
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
	out := &SetLogFilePathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAPITracingDataCollectorClient) GetIncludeModules(ctx context.Context, in *GetIncludeModulesRequest, opts ...dcerpc.CallOption) (*GetIncludeModulesResponse, error) {
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
	out := &GetIncludeModulesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAPITracingDataCollectorClient) SetIncludeModules(ctx context.Context, in *SetIncludeModulesRequest, opts ...dcerpc.CallOption) (*SetIncludeModulesResponse, error) {
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
	out := &SetIncludeModulesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAPITracingDataCollectorClient) GetIncludeAPIs(ctx context.Context, in *GetIncludeAPIsRequest, opts ...dcerpc.CallOption) (*GetIncludeAPIsResponse, error) {
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
	out := &GetIncludeAPIsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAPITracingDataCollectorClient) SetIncludeAPIs(ctx context.Context, in *SetIncludeAPIsRequest, opts ...dcerpc.CallOption) (*SetIncludeAPIsResponse, error) {
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
	out := &SetIncludeAPIsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAPITracingDataCollectorClient) GetExcludeAPIs(ctx context.Context, in *GetExcludeAPIsRequest, opts ...dcerpc.CallOption) (*GetExcludeAPIsResponse, error) {
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
	out := &GetExcludeAPIsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAPITracingDataCollectorClient) SetExcludeAPIs(ctx context.Context, in *SetExcludeAPIsRequest, opts ...dcerpc.CallOption) (*SetExcludeAPIsResponse, error) {
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
	out := &SetExcludeAPIsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAPITracingDataCollectorClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAPITracingDataCollectorClient) IPID(ctx context.Context, ipid *dcom.IPID) APITracingDataCollectorClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAPITracingDataCollectorClient{
		DataCollectorClient: o.DataCollectorClient.IPID(ctx, ipid),
		cc:                  o.cc,
		ipid:                ipid,
	}
}
func NewAPITracingDataCollectorClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (APITracingDataCollectorClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(APITracingDataCollectorSyntaxV0_0))...)
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
	return &xxx_DefaultAPITracingDataCollectorClient{
		DataCollectorClient: base,
		cc:                  cc,
		ipid:                ipid,
	}, nil
}

// xxx_GetLogAPINamesOnlyOperation structure represents the LogApiNamesOnly operation
type xxx_GetLogAPINamesOnlyOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	LogAPINames int16          `idl:"name:logapinames" json:"log_api_names"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLogAPINamesOnlyOperation) OpNum() int { return 32 }

func (o *xxx_GetLogAPINamesOnlyOperation) OpName() string {
	return "/IApiTracingDataCollector/v0/LogApiNamesOnly"
}

func (o *xxx_GetLogAPINamesOnlyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogAPINamesOnlyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLogAPINamesOnlyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLogAPINamesOnlyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogAPINamesOnlyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// logapinames {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.LogAPINames); err != nil {
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

func (o *xxx_GetLogAPINamesOnlyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// logapinames {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.LogAPINames); err != nil {
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

// GetLogAPINamesOnlyRequest structure represents the LogApiNamesOnly operation request
type GetLogAPINamesOnlyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLogAPINamesOnlyRequest) xxx_ToOp(ctx context.Context) *xxx_GetLogAPINamesOnlyOperation {
	if o == nil {
		return &xxx_GetLogAPINamesOnlyOperation{}
	}
	return &xxx_GetLogAPINamesOnlyOperation{
		This: o.This,
	}
}

func (o *GetLogAPINamesOnlyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLogAPINamesOnlyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLogAPINamesOnlyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetLogAPINamesOnlyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLogAPINamesOnlyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLogAPINamesOnlyResponse structure represents the LogApiNamesOnly operation response
type GetLogAPINamesOnlyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	LogAPINames int16          `idl:"name:logapinames" json:"log_api_names"`
	// Return: The LogApiNamesOnly return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLogAPINamesOnlyResponse) xxx_ToOp(ctx context.Context) *xxx_GetLogAPINamesOnlyOperation {
	if o == nil {
		return &xxx_GetLogAPINamesOnlyOperation{}
	}
	return &xxx_GetLogAPINamesOnlyOperation{
		That:        o.That,
		LogAPINames: o.LogAPINames,
		Return:      o.Return,
	}
}

func (o *GetLogAPINamesOnlyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLogAPINamesOnlyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LogAPINames = op.LogAPINames
	o.Return = op.Return
}
func (o *GetLogAPINamesOnlyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetLogAPINamesOnlyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLogAPINamesOnlyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetLogAPINamesOnlyOperation structure represents the LogApiNamesOnly operation
type xxx_SetLogAPINamesOnlyOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	LogAPINames int16          `idl:"name:logapinames" json:"log_api_names"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetLogAPINamesOnlyOperation) OpNum() int { return 33 }

func (o *xxx_SetLogAPINamesOnlyOperation) OpName() string {
	return "/IApiTracingDataCollector/v0/LogApiNamesOnly"
}

func (o *xxx_SetLogAPINamesOnlyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogAPINamesOnlyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// logapinames {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.LogAPINames); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogAPINamesOnlyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// logapinames {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.LogAPINames); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogAPINamesOnlyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogAPINamesOnlyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetLogAPINamesOnlyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetLogAPINamesOnlyRequest structure represents the LogApiNamesOnly operation request
type SetLogAPINamesOnlyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	LogAPINames int16          `idl:"name:logapinames" json:"log_api_names"`
}

func (o *SetLogAPINamesOnlyRequest) xxx_ToOp(ctx context.Context) *xxx_SetLogAPINamesOnlyOperation {
	if o == nil {
		return &xxx_SetLogAPINamesOnlyOperation{}
	}
	return &xxx_SetLogAPINamesOnlyOperation{
		This:        o.This,
		LogAPINames: o.LogAPINames,
	}
}

func (o *SetLogAPINamesOnlyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetLogAPINamesOnlyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LogAPINames = op.LogAPINames
}
func (o *SetLogAPINamesOnlyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetLogAPINamesOnlyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLogAPINamesOnlyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetLogAPINamesOnlyResponse structure represents the LogApiNamesOnly operation response
type SetLogAPINamesOnlyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The LogApiNamesOnly return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetLogAPINamesOnlyResponse) xxx_ToOp(ctx context.Context) *xxx_SetLogAPINamesOnlyOperation {
	if o == nil {
		return &xxx_SetLogAPINamesOnlyOperation{}
	}
	return &xxx_SetLogAPINamesOnlyOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetLogAPINamesOnlyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetLogAPINamesOnlyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetLogAPINamesOnlyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetLogAPINamesOnlyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLogAPINamesOnlyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLogAPIsRecursivelyOperation structure represents the LogApisRecursively operation
type xxx_GetLogAPIsRecursivelyOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	LogRecursively int16          `idl:"name:logrecursively" json:"log_recursively"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLogAPIsRecursivelyOperation) OpNum() int { return 34 }

func (o *xxx_GetLogAPIsRecursivelyOperation) OpName() string {
	return "/IApiTracingDataCollector/v0/LogApisRecursively"
}

func (o *xxx_GetLogAPIsRecursivelyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogAPIsRecursivelyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLogAPIsRecursivelyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLogAPIsRecursivelyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogAPIsRecursivelyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// logrecursively {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.LogRecursively); err != nil {
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

func (o *xxx_GetLogAPIsRecursivelyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// logrecursively {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.LogRecursively); err != nil {
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

// GetLogAPIsRecursivelyRequest structure represents the LogApisRecursively operation request
type GetLogAPIsRecursivelyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLogAPIsRecursivelyRequest) xxx_ToOp(ctx context.Context) *xxx_GetLogAPIsRecursivelyOperation {
	if o == nil {
		return &xxx_GetLogAPIsRecursivelyOperation{}
	}
	return &xxx_GetLogAPIsRecursivelyOperation{
		This: o.This,
	}
}

func (o *GetLogAPIsRecursivelyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLogAPIsRecursivelyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLogAPIsRecursivelyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetLogAPIsRecursivelyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLogAPIsRecursivelyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLogAPIsRecursivelyResponse structure represents the LogApisRecursively operation response
type GetLogAPIsRecursivelyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	LogRecursively int16          `idl:"name:logrecursively" json:"log_recursively"`
	// Return: The LogApisRecursively return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLogAPIsRecursivelyResponse) xxx_ToOp(ctx context.Context) *xxx_GetLogAPIsRecursivelyOperation {
	if o == nil {
		return &xxx_GetLogAPIsRecursivelyOperation{}
	}
	return &xxx_GetLogAPIsRecursivelyOperation{
		That:           o.That,
		LogRecursively: o.LogRecursively,
		Return:         o.Return,
	}
}

func (o *GetLogAPIsRecursivelyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLogAPIsRecursivelyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LogRecursively = op.LogRecursively
	o.Return = op.Return
}
func (o *GetLogAPIsRecursivelyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetLogAPIsRecursivelyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLogAPIsRecursivelyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetLogAPIsRecursivelyOperation structure represents the LogApisRecursively operation
type xxx_SetLogAPIsRecursivelyOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	LogRecursively int16          `idl:"name:logrecursively" json:"log_recursively"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetLogAPIsRecursivelyOperation) OpNum() int { return 35 }

func (o *xxx_SetLogAPIsRecursivelyOperation) OpName() string {
	return "/IApiTracingDataCollector/v0/LogApisRecursively"
}

func (o *xxx_SetLogAPIsRecursivelyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogAPIsRecursivelyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// logrecursively {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.LogRecursively); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogAPIsRecursivelyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// logrecursively {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.LogRecursively); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogAPIsRecursivelyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogAPIsRecursivelyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetLogAPIsRecursivelyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetLogAPIsRecursivelyRequest structure represents the LogApisRecursively operation request
type SetLogAPIsRecursivelyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	LogRecursively int16          `idl:"name:logrecursively" json:"log_recursively"`
}

func (o *SetLogAPIsRecursivelyRequest) xxx_ToOp(ctx context.Context) *xxx_SetLogAPIsRecursivelyOperation {
	if o == nil {
		return &xxx_SetLogAPIsRecursivelyOperation{}
	}
	return &xxx_SetLogAPIsRecursivelyOperation{
		This:           o.This,
		LogRecursively: o.LogRecursively,
	}
}

func (o *SetLogAPIsRecursivelyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetLogAPIsRecursivelyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LogRecursively = op.LogRecursively
}
func (o *SetLogAPIsRecursivelyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetLogAPIsRecursivelyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLogAPIsRecursivelyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetLogAPIsRecursivelyResponse structure represents the LogApisRecursively operation response
type SetLogAPIsRecursivelyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The LogApisRecursively return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetLogAPIsRecursivelyResponse) xxx_ToOp(ctx context.Context) *xxx_SetLogAPIsRecursivelyOperation {
	if o == nil {
		return &xxx_SetLogAPIsRecursivelyOperation{}
	}
	return &xxx_SetLogAPIsRecursivelyOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetLogAPIsRecursivelyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetLogAPIsRecursivelyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetLogAPIsRecursivelyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetLogAPIsRecursivelyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLogAPIsRecursivelyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetExePathOperation structure represents the ExePath operation
type xxx_GetExePathOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	ExePath *oaut.String   `idl:"name:exepath" json:"exe_path"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetExePathOperation) OpNum() int { return 36 }

func (o *xxx_GetExePathOperation) OpName() string { return "/IApiTracingDataCollector/v0/ExePath" }

func (o *xxx_GetExePathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExePathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetExePathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetExePathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExePathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// exepath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ExePath != nil {
			_ptr_exepath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ExePath != nil {
					if err := o.ExePath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExePath, _ptr_exepath); err != nil {
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

func (o *xxx_GetExePathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// exepath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_exepath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ExePath == nil {
				o.ExePath = &oaut.String{}
			}
			if err := o.ExePath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_exepath := func(ptr interface{}) { o.ExePath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ExePath, _s_exepath, _ptr_exepath); err != nil {
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

// GetExePathRequest structure represents the ExePath operation request
type GetExePathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetExePathRequest) xxx_ToOp(ctx context.Context) *xxx_GetExePathOperation {
	if o == nil {
		return &xxx_GetExePathOperation{}
	}
	return &xxx_GetExePathOperation{
		This: o.This,
	}
}

func (o *GetExePathRequest) xxx_FromOp(ctx context.Context, op *xxx_GetExePathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetExePathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetExePathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExePathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetExePathResponse structure represents the ExePath operation response
type GetExePathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	ExePath *oaut.String   `idl:"name:exepath" json:"exe_path"`
	// Return: The ExePath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetExePathResponse) xxx_ToOp(ctx context.Context) *xxx_GetExePathOperation {
	if o == nil {
		return &xxx_GetExePathOperation{}
	}
	return &xxx_GetExePathOperation{
		That:    o.That,
		ExePath: o.ExePath,
		Return:  o.Return,
	}
}

func (o *GetExePathResponse) xxx_FromOp(ctx context.Context, op *xxx_GetExePathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ExePath = op.ExePath
	o.Return = op.Return
}
func (o *GetExePathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetExePathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExePathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetExePathOperation structure represents the ExePath operation
type xxx_SetExePathOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	ExePath *oaut.String   `idl:"name:exepath" json:"exe_path"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetExePathOperation) OpNum() int { return 37 }

func (o *xxx_SetExePathOperation) OpName() string { return "/IApiTracingDataCollector/v0/ExePath" }

func (o *xxx_SetExePathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExePathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// exepath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ExePath != nil {
			_ptr_exepath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ExePath != nil {
					if err := o.ExePath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExePath, _ptr_exepath); err != nil {
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

func (o *xxx_SetExePathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// exepath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_exepath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ExePath == nil {
				o.ExePath = &oaut.String{}
			}
			if err := o.ExePath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_exepath := func(ptr interface{}) { o.ExePath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ExePath, _s_exepath, _ptr_exepath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExePathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExePathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetExePathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetExePathRequest structure represents the ExePath operation request
type SetExePathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	ExePath *oaut.String   `idl:"name:exepath" json:"exe_path"`
}

func (o *SetExePathRequest) xxx_ToOp(ctx context.Context) *xxx_SetExePathOperation {
	if o == nil {
		return &xxx_SetExePathOperation{}
	}
	return &xxx_SetExePathOperation{
		This:    o.This,
		ExePath: o.ExePath,
	}
}

func (o *SetExePathRequest) xxx_FromOp(ctx context.Context, op *xxx_SetExePathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ExePath = op.ExePath
}
func (o *SetExePathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetExePathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExePathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetExePathResponse structure represents the ExePath operation response
type SetExePathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ExePath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetExePathResponse) xxx_ToOp(ctx context.Context) *xxx_SetExePathOperation {
	if o == nil {
		return &xxx_SetExePathOperation{}
	}
	return &xxx_SetExePathOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetExePathResponse) xxx_FromOp(ctx context.Context, op *xxx_SetExePathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetExePathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetExePathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExePathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLogFilePathOperation structure represents the LogFilePath operation
type xxx_GetLogFilePathOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Logfilepath *oaut.String   `idl:"name:logfilepath" json:"logfilepath"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLogFilePathOperation) OpNum() int { return 38 }

func (o *xxx_GetLogFilePathOperation) OpName() string {
	return "/IApiTracingDataCollector/v0/LogFilePath"
}

func (o *xxx_GetLogFilePathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogFilePathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLogFilePathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLogFilePathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogFilePathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// logfilepath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Logfilepath != nil {
			_ptr_logfilepath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Logfilepath != nil {
					if err := o.Logfilepath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Logfilepath, _ptr_logfilepath); err != nil {
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

func (o *xxx_GetLogFilePathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// logfilepath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_logfilepath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Logfilepath == nil {
				o.Logfilepath = &oaut.String{}
			}
			if err := o.Logfilepath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_logfilepath := func(ptr interface{}) { o.Logfilepath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Logfilepath, _s_logfilepath, _ptr_logfilepath); err != nil {
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

// GetLogFilePathRequest structure represents the LogFilePath operation request
type GetLogFilePathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLogFilePathRequest) xxx_ToOp(ctx context.Context) *xxx_GetLogFilePathOperation {
	if o == nil {
		return &xxx_GetLogFilePathOperation{}
	}
	return &xxx_GetLogFilePathOperation{
		This: o.This,
	}
}

func (o *GetLogFilePathRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLogFilePathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLogFilePathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetLogFilePathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLogFilePathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLogFilePathResponse structure represents the LogFilePath operation response
type GetLogFilePathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Logfilepath *oaut.String   `idl:"name:logfilepath" json:"logfilepath"`
	// Return: The LogFilePath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLogFilePathResponse) xxx_ToOp(ctx context.Context) *xxx_GetLogFilePathOperation {
	if o == nil {
		return &xxx_GetLogFilePathOperation{}
	}
	return &xxx_GetLogFilePathOperation{
		That:        o.That,
		Logfilepath: o.Logfilepath,
		Return:      o.Return,
	}
}

func (o *GetLogFilePathResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLogFilePathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Logfilepath = op.Logfilepath
	o.Return = op.Return
}
func (o *GetLogFilePathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetLogFilePathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLogFilePathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetLogFilePathOperation structure represents the LogFilePath operation
type xxx_SetLogFilePathOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Logfilepath *oaut.String   `idl:"name:logfilepath" json:"logfilepath"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetLogFilePathOperation) OpNum() int { return 39 }

func (o *xxx_SetLogFilePathOperation) OpName() string {
	return "/IApiTracingDataCollector/v0/LogFilePath"
}

func (o *xxx_SetLogFilePathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogFilePathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// logfilepath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Logfilepath != nil {
			_ptr_logfilepath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Logfilepath != nil {
					if err := o.Logfilepath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Logfilepath, _ptr_logfilepath); err != nil {
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

func (o *xxx_SetLogFilePathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// logfilepath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_logfilepath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Logfilepath == nil {
				o.Logfilepath = &oaut.String{}
			}
			if err := o.Logfilepath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_logfilepath := func(ptr interface{}) { o.Logfilepath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Logfilepath, _s_logfilepath, _ptr_logfilepath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogFilePathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogFilePathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetLogFilePathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetLogFilePathRequest structure represents the LogFilePath operation request
type SetLogFilePathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	Logfilepath *oaut.String   `idl:"name:logfilepath" json:"logfilepath"`
}

func (o *SetLogFilePathRequest) xxx_ToOp(ctx context.Context) *xxx_SetLogFilePathOperation {
	if o == nil {
		return &xxx_SetLogFilePathOperation{}
	}
	return &xxx_SetLogFilePathOperation{
		This:        o.This,
		Logfilepath: o.Logfilepath,
	}
}

func (o *SetLogFilePathRequest) xxx_FromOp(ctx context.Context, op *xxx_SetLogFilePathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Logfilepath = op.Logfilepath
}
func (o *SetLogFilePathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetLogFilePathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLogFilePathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetLogFilePathResponse structure represents the LogFilePath operation response
type SetLogFilePathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The LogFilePath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetLogFilePathResponse) xxx_ToOp(ctx context.Context) *xxx_SetLogFilePathOperation {
	if o == nil {
		return &xxx_SetLogFilePathOperation{}
	}
	return &xxx_SetLogFilePathOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetLogFilePathResponse) xxx_FromOp(ctx context.Context, op *xxx_SetLogFilePathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetLogFilePathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetLogFilePathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLogFilePathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIncludeModulesOperation structure represents the IncludeModules operation
type xxx_GetIncludeModulesOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	IncludeModules *oaut.SafeArray `idl:"name:includemodules" json:"include_modules"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIncludeModulesOperation) OpNum() int { return 40 }

func (o *xxx_GetIncludeModulesOperation) OpName() string {
	return "/IApiTracingDataCollector/v0/IncludeModules"
}

func (o *xxx_GetIncludeModulesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIncludeModulesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIncludeModulesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIncludeModulesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIncludeModulesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// includemodules {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.IncludeModules != nil {
			_ptr_includemodules := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.IncludeModules != nil {
					if err := o.IncludeModules.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.IncludeModules, _ptr_includemodules); err != nil {
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

func (o *xxx_GetIncludeModulesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// includemodules {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_includemodules := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.IncludeModules == nil {
				o.IncludeModules = &oaut.SafeArray{}
			}
			if err := o.IncludeModules.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_includemodules := func(ptr interface{}) { o.IncludeModules = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.IncludeModules, _s_includemodules, _ptr_includemodules); err != nil {
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

// GetIncludeModulesRequest structure represents the IncludeModules operation request
type GetIncludeModulesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIncludeModulesRequest) xxx_ToOp(ctx context.Context) *xxx_GetIncludeModulesOperation {
	if o == nil {
		return &xxx_GetIncludeModulesOperation{}
	}
	return &xxx_GetIncludeModulesOperation{
		This: o.This,
	}
}

func (o *GetIncludeModulesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIncludeModulesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIncludeModulesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetIncludeModulesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIncludeModulesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIncludeModulesResponse structure represents the IncludeModules operation response
type GetIncludeModulesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	IncludeModules *oaut.SafeArray `idl:"name:includemodules" json:"include_modules"`
	// Return: The IncludeModules return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIncludeModulesResponse) xxx_ToOp(ctx context.Context) *xxx_GetIncludeModulesOperation {
	if o == nil {
		return &xxx_GetIncludeModulesOperation{}
	}
	return &xxx_GetIncludeModulesOperation{
		That:           o.That,
		IncludeModules: o.IncludeModules,
		Return:         o.Return,
	}
}

func (o *GetIncludeModulesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIncludeModulesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IncludeModules = op.IncludeModules
	o.Return = op.Return
}
func (o *GetIncludeModulesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetIncludeModulesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIncludeModulesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetIncludeModulesOperation structure represents the IncludeModules operation
type xxx_SetIncludeModulesOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	IncludeModules *oaut.SafeArray `idl:"name:includemodules" json:"include_modules"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetIncludeModulesOperation) OpNum() int { return 41 }

func (o *xxx_SetIncludeModulesOperation) OpName() string {
	return "/IApiTracingDataCollector/v0/IncludeModules"
}

func (o *xxx_SetIncludeModulesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIncludeModulesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// includemodules {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.IncludeModules != nil {
			_ptr_includemodules := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.IncludeModules != nil {
					if err := o.IncludeModules.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.IncludeModules, _ptr_includemodules); err != nil {
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

func (o *xxx_SetIncludeModulesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// includemodules {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_includemodules := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.IncludeModules == nil {
				o.IncludeModules = &oaut.SafeArray{}
			}
			if err := o.IncludeModules.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_includemodules := func(ptr interface{}) { o.IncludeModules = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.IncludeModules, _s_includemodules, _ptr_includemodules); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIncludeModulesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIncludeModulesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetIncludeModulesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetIncludeModulesRequest structure represents the IncludeModules operation request
type SetIncludeModulesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	IncludeModules *oaut.SafeArray `idl:"name:includemodules" json:"include_modules"`
}

func (o *SetIncludeModulesRequest) xxx_ToOp(ctx context.Context) *xxx_SetIncludeModulesOperation {
	if o == nil {
		return &xxx_SetIncludeModulesOperation{}
	}
	return &xxx_SetIncludeModulesOperation{
		This:           o.This,
		IncludeModules: o.IncludeModules,
	}
}

func (o *SetIncludeModulesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetIncludeModulesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.IncludeModules = op.IncludeModules
}
func (o *SetIncludeModulesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetIncludeModulesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetIncludeModulesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetIncludeModulesResponse structure represents the IncludeModules operation response
type SetIncludeModulesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The IncludeModules return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetIncludeModulesResponse) xxx_ToOp(ctx context.Context) *xxx_SetIncludeModulesOperation {
	if o == nil {
		return &xxx_SetIncludeModulesOperation{}
	}
	return &xxx_SetIncludeModulesOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetIncludeModulesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetIncludeModulesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetIncludeModulesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetIncludeModulesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetIncludeModulesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIncludeAPIsOperation structure represents the IncludeApis operation
type xxx_GetIncludeAPIsOperation struct {
	This        *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat  `idl:"name:That" json:"that"`
	IncludeAPIs *oaut.SafeArray `idl:"name:includeapis" json:"include_apis"`
	Return      int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIncludeAPIsOperation) OpNum() int { return 42 }

func (o *xxx_GetIncludeAPIsOperation) OpName() string {
	return "/IApiTracingDataCollector/v0/IncludeApis"
}

func (o *xxx_GetIncludeAPIsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIncludeAPIsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIncludeAPIsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIncludeAPIsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIncludeAPIsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// includeapis {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.IncludeAPIs != nil {
			_ptr_includeapis := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.IncludeAPIs != nil {
					if err := o.IncludeAPIs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.IncludeAPIs, _ptr_includeapis); err != nil {
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

func (o *xxx_GetIncludeAPIsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// includeapis {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_includeapis := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.IncludeAPIs == nil {
				o.IncludeAPIs = &oaut.SafeArray{}
			}
			if err := o.IncludeAPIs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_includeapis := func(ptr interface{}) { o.IncludeAPIs = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.IncludeAPIs, _s_includeapis, _ptr_includeapis); err != nil {
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

// GetIncludeAPIsRequest structure represents the IncludeApis operation request
type GetIncludeAPIsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIncludeAPIsRequest) xxx_ToOp(ctx context.Context) *xxx_GetIncludeAPIsOperation {
	if o == nil {
		return &xxx_GetIncludeAPIsOperation{}
	}
	return &xxx_GetIncludeAPIsOperation{
		This: o.This,
	}
}

func (o *GetIncludeAPIsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIncludeAPIsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIncludeAPIsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetIncludeAPIsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIncludeAPIsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIncludeAPIsResponse structure represents the IncludeApis operation response
type GetIncludeAPIsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat  `idl:"name:That" json:"that"`
	IncludeAPIs *oaut.SafeArray `idl:"name:includeapis" json:"include_apis"`
	// Return: The IncludeApis return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIncludeAPIsResponse) xxx_ToOp(ctx context.Context) *xxx_GetIncludeAPIsOperation {
	if o == nil {
		return &xxx_GetIncludeAPIsOperation{}
	}
	return &xxx_GetIncludeAPIsOperation{
		That:        o.That,
		IncludeAPIs: o.IncludeAPIs,
		Return:      o.Return,
	}
}

func (o *GetIncludeAPIsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIncludeAPIsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IncludeAPIs = op.IncludeAPIs
	o.Return = op.Return
}
func (o *GetIncludeAPIsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetIncludeAPIsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIncludeAPIsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetIncludeAPIsOperation structure represents the IncludeApis operation
type xxx_SetIncludeAPIsOperation struct {
	This        *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat  `idl:"name:That" json:"that"`
	IncludeAPIs *oaut.SafeArray `idl:"name:includeapis" json:"include_apis"`
	Return      int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetIncludeAPIsOperation) OpNum() int { return 43 }

func (o *xxx_SetIncludeAPIsOperation) OpName() string {
	return "/IApiTracingDataCollector/v0/IncludeApis"
}

func (o *xxx_SetIncludeAPIsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIncludeAPIsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// includeapis {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.IncludeAPIs != nil {
			_ptr_includeapis := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.IncludeAPIs != nil {
					if err := o.IncludeAPIs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.IncludeAPIs, _ptr_includeapis); err != nil {
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

func (o *xxx_SetIncludeAPIsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// includeapis {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_includeapis := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.IncludeAPIs == nil {
				o.IncludeAPIs = &oaut.SafeArray{}
			}
			if err := o.IncludeAPIs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_includeapis := func(ptr interface{}) { o.IncludeAPIs = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.IncludeAPIs, _s_includeapis, _ptr_includeapis); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIncludeAPIsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIncludeAPIsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetIncludeAPIsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetIncludeAPIsRequest structure represents the IncludeApis operation request
type SetIncludeAPIsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis  `idl:"name:This" json:"this"`
	IncludeAPIs *oaut.SafeArray `idl:"name:includeapis" json:"include_apis"`
}

func (o *SetIncludeAPIsRequest) xxx_ToOp(ctx context.Context) *xxx_SetIncludeAPIsOperation {
	if o == nil {
		return &xxx_SetIncludeAPIsOperation{}
	}
	return &xxx_SetIncludeAPIsOperation{
		This:        o.This,
		IncludeAPIs: o.IncludeAPIs,
	}
}

func (o *SetIncludeAPIsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetIncludeAPIsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.IncludeAPIs = op.IncludeAPIs
}
func (o *SetIncludeAPIsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetIncludeAPIsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetIncludeAPIsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetIncludeAPIsResponse structure represents the IncludeApis operation response
type SetIncludeAPIsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The IncludeApis return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetIncludeAPIsResponse) xxx_ToOp(ctx context.Context) *xxx_SetIncludeAPIsOperation {
	if o == nil {
		return &xxx_SetIncludeAPIsOperation{}
	}
	return &xxx_SetIncludeAPIsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetIncludeAPIsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetIncludeAPIsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetIncludeAPIsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetIncludeAPIsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetIncludeAPIsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetExcludeAPIsOperation structure represents the ExcludeApis operation
type xxx_GetExcludeAPIsOperation struct {
	This        *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ExcludeAPIs *oaut.SafeArray `idl:"name:excludeapis" json:"exclude_apis"`
	Return      int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetExcludeAPIsOperation) OpNum() int { return 44 }

func (o *xxx_GetExcludeAPIsOperation) OpName() string {
	return "/IApiTracingDataCollector/v0/ExcludeApis"
}

func (o *xxx_GetExcludeAPIsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExcludeAPIsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetExcludeAPIsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetExcludeAPIsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExcludeAPIsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// excludeapis {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.ExcludeAPIs != nil {
			_ptr_excludeapis := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ExcludeAPIs != nil {
					if err := o.ExcludeAPIs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExcludeAPIs, _ptr_excludeapis); err != nil {
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

func (o *xxx_GetExcludeAPIsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// excludeapis {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_excludeapis := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ExcludeAPIs == nil {
				o.ExcludeAPIs = &oaut.SafeArray{}
			}
			if err := o.ExcludeAPIs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_excludeapis := func(ptr interface{}) { o.ExcludeAPIs = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.ExcludeAPIs, _s_excludeapis, _ptr_excludeapis); err != nil {
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

// GetExcludeAPIsRequest structure represents the ExcludeApis operation request
type GetExcludeAPIsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetExcludeAPIsRequest) xxx_ToOp(ctx context.Context) *xxx_GetExcludeAPIsOperation {
	if o == nil {
		return &xxx_GetExcludeAPIsOperation{}
	}
	return &xxx_GetExcludeAPIsOperation{
		This: o.This,
	}
}

func (o *GetExcludeAPIsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetExcludeAPIsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetExcludeAPIsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetExcludeAPIsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExcludeAPIsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetExcludeAPIsResponse structure represents the ExcludeApis operation response
type GetExcludeAPIsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ExcludeAPIs *oaut.SafeArray `idl:"name:excludeapis" json:"exclude_apis"`
	// Return: The ExcludeApis return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetExcludeAPIsResponse) xxx_ToOp(ctx context.Context) *xxx_GetExcludeAPIsOperation {
	if o == nil {
		return &xxx_GetExcludeAPIsOperation{}
	}
	return &xxx_GetExcludeAPIsOperation{
		That:        o.That,
		ExcludeAPIs: o.ExcludeAPIs,
		Return:      o.Return,
	}
}

func (o *GetExcludeAPIsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetExcludeAPIsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ExcludeAPIs = op.ExcludeAPIs
	o.Return = op.Return
}
func (o *GetExcludeAPIsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetExcludeAPIsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExcludeAPIsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetExcludeAPIsOperation structure represents the ExcludeApis operation
type xxx_SetExcludeAPIsOperation struct {
	This        *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ExcludeAPIs *oaut.SafeArray `idl:"name:excludeapis" json:"exclude_apis"`
	Return      int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetExcludeAPIsOperation) OpNum() int { return 45 }

func (o *xxx_SetExcludeAPIsOperation) OpName() string {
	return "/IApiTracingDataCollector/v0/ExcludeApis"
}

func (o *xxx_SetExcludeAPIsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExcludeAPIsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// excludeapis {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.ExcludeAPIs != nil {
			_ptr_excludeapis := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ExcludeAPIs != nil {
					if err := o.ExcludeAPIs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExcludeAPIs, _ptr_excludeapis); err != nil {
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

func (o *xxx_SetExcludeAPIsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// excludeapis {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_excludeapis := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ExcludeAPIs == nil {
				o.ExcludeAPIs = &oaut.SafeArray{}
			}
			if err := o.ExcludeAPIs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_excludeapis := func(ptr interface{}) { o.ExcludeAPIs = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.ExcludeAPIs, _s_excludeapis, _ptr_excludeapis); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExcludeAPIsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExcludeAPIsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetExcludeAPIsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetExcludeAPIsRequest structure represents the ExcludeApis operation request
type SetExcludeAPIsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis  `idl:"name:This" json:"this"`
	ExcludeAPIs *oaut.SafeArray `idl:"name:excludeapis" json:"exclude_apis"`
}

func (o *SetExcludeAPIsRequest) xxx_ToOp(ctx context.Context) *xxx_SetExcludeAPIsOperation {
	if o == nil {
		return &xxx_SetExcludeAPIsOperation{}
	}
	return &xxx_SetExcludeAPIsOperation{
		This:        o.This,
		ExcludeAPIs: o.ExcludeAPIs,
	}
}

func (o *SetExcludeAPIsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetExcludeAPIsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ExcludeAPIs = op.ExcludeAPIs
}
func (o *SetExcludeAPIsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetExcludeAPIsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExcludeAPIsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetExcludeAPIsResponse structure represents the ExcludeApis operation response
type SetExcludeAPIsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ExcludeApis return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetExcludeAPIsResponse) xxx_ToOp(ctx context.Context) *xxx_SetExcludeAPIsOperation {
	if o == nil {
		return &xxx_SetExcludeAPIsOperation{}
	}
	return &xxx_SetExcludeAPIsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetExcludeAPIsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetExcludeAPIsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetExcludeAPIsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetExcludeAPIsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExcludeAPIsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
