package ivdsasync

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	vds "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds"
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
	_ = iunknown.GoPackage
	_ = vds.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/vds"
)

var (
	// IVdsAsync interface identifier d5d23b6d-5a55-4492-9889-397a3c2d2dbc
	AsyncIID = &dcom.IID{Data1: 0xd5d23b6d, Data2: 0x5a55, Data3: 0x4492, Data4: []byte{0x98, 0x89, 0x39, 0x7a, 0x3c, 0x2d, 0x2d, 0xbc}}
	// Syntax UUID
	AsyncSyntaxUUID = &uuid.UUID{TimeLow: 0xd5d23b6d, TimeMid: 0x5a55, TimeHiAndVersion: 0x4492, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x89, Node: [6]uint8{0x39, 0x7a, 0x3c, 0x2d, 0x2d, 0xbc}}
	// Syntax ID
	AsyncSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AsyncSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsAsync interface.
type AsyncClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The Cancel method cancels the asynchronous operation.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Cancel(context.Context, *CancelRequest, ...dcerpc.CallOption) (*CancelResponse, error)

	// The Wait method blocks and returns when the asynchronous operation has either finished
	// successfully or failed.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Wait(context.Context, *WaitRequest, ...dcerpc.CallOption) (*WaitResponse, error)

	// The QueryStatus method retrieves the status of the asynchronous operation.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryStatus(context.Context, *QueryStatusRequest, ...dcerpc.CallOption) (*QueryStatusResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AsyncClient
}

type xxx_DefaultAsyncClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAsyncClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAsyncClient) Cancel(ctx context.Context, in *CancelRequest, opts ...dcerpc.CallOption) (*CancelResponse, error) {
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
	out := &CancelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAsyncClient) Wait(ctx context.Context, in *WaitRequest, opts ...dcerpc.CallOption) (*WaitResponse, error) {
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
	out := &WaitResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAsyncClient) QueryStatus(ctx context.Context, in *QueryStatusRequest, opts ...dcerpc.CallOption) (*QueryStatusResponse, error) {
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
	out := &QueryStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAsyncClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAsyncClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAsyncClient) IPID(ctx context.Context, ipid *dcom.IPID) AsyncClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAsyncClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewAsyncClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AsyncClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AsyncSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultAsyncClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_CancelOperation structure represents the Cancel operation
type xxx_CancelOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CancelOperation) OpNum() int { return 3 }

func (o *xxx_CancelOperation) OpName() string { return "/IVdsAsync/v0/Cancel" }

func (o *xxx_CancelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CancelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CancelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CancelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CancelRequest structure represents the Cancel operation request
type CancelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CancelRequest) xxx_ToOp(ctx context.Context, op *xxx_CancelOperation) *xxx_CancelOperation {
	if op == nil {
		op = &xxx_CancelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CancelRequest) xxx_FromOp(ctx context.Context, op *xxx_CancelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CancelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CancelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CancelResponse structure represents the Cancel operation response
type CancelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Cancel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CancelResponse) xxx_ToOp(ctx context.Context, op *xxx_CancelOperation) *xxx_CancelOperation {
	if op == nil {
		op = &xxx_CancelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CancelResponse) xxx_FromOp(ctx context.Context, op *xxx_CancelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CancelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CancelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WaitOperation structure represents the Wait operation
type xxx_WaitOperation struct {
	This        *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat   `idl:"name:That" json:"that"`
	HResult     int32            `idl:"name:pHrResult" json:"hresult"`
	AsyncOutput *vds.AsyncOutput `idl:"name:pAsyncOut" json:"async_output"`
	Return      int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_WaitOperation) OpNum() int { return 4 }

func (o *xxx_WaitOperation) OpName() string { return "/IVdsAsync/v0/Wait" }

func (o *xxx_WaitOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_WaitOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_WaitOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pHrResult {out} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.HResult); err != nil {
			return err
		}
	}
	// pAsyncOut {out} (1:{pointer=ref}*(1))(2:{alias=VDS_ASYNC_OUTPUT}(struct))
	{
		if o.AsyncOutput != nil {
			if err := o.AsyncOutput.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.AsyncOutput{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_WaitOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pHrResult {out} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.HResult); err != nil {
			return err
		}
	}
	// pAsyncOut {out} (1:{pointer=ref}*(1))(2:{alias=VDS_ASYNC_OUTPUT}(struct))
	{
		if o.AsyncOutput == nil {
			o.AsyncOutput = &vds.AsyncOutput{}
		}
		if err := o.AsyncOutput.UnmarshalNDR(ctx, w); err != nil {
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

// WaitRequest structure represents the Wait operation request
type WaitRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *WaitRequest) xxx_ToOp(ctx context.Context, op *xxx_WaitOperation) *xxx_WaitOperation {
	if op == nil {
		op = &xxx_WaitOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *WaitRequest) xxx_FromOp(ctx context.Context, op *xxx_WaitOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *WaitRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WaitRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WaitResponse structure represents the Wait operation response
type WaitResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pHrResult: A pointer to a variable which, if the Wait method successfully completes,
	// receives the returned HRESULT.
	HResult int32 `idl:"name:pHrResult" json:"hresult"`
	// pAsyncOut: A pointer to a VDS_ASYNC_OUTPUT structure that, if the asynchronous operation
	// is successfully completed, receives extra information about the operation, if any
	// information exists. Multiple methods from other interfaces also return async objects.
	// Consult the method that returned the async object to determine what extra information
	// to return, if any. If the asynchronous operation fails, pAsyncOut MAY be left as
	// is without returning any value.
	AsyncOutput *vds.AsyncOutput `idl:"name:pAsyncOut" json:"async_output"`
	// Return: The Wait return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WaitResponse) xxx_ToOp(ctx context.Context, op *xxx_WaitOperation) *xxx_WaitOperation {
	if op == nil {
		op = &xxx_WaitOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.HResult = o.HResult
	op.AsyncOutput = o.AsyncOutput
	op.Return = o.Return
	return op
}

func (o *WaitResponse) xxx_FromOp(ctx context.Context, op *xxx_WaitOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.HResult = op.HResult
	o.AsyncOutput = op.AsyncOutput
	o.Return = op.Return
}
func (o *WaitResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WaitResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryStatusOperation structure represents the QueryStatus operation
type xxx_QueryStatusOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	HResult          int32          `idl:"name:pHrResult" json:"hresult"`
	PercentCompleted uint32         `idl:"name:pulPercentCompleted" json:"percent_completed"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryStatusOperation) OpNum() int { return 5 }

func (o *xxx_QueryStatusOperation) OpName() string { return "/IVdsAsync/v0/QueryStatus" }

func (o *xxx_QueryStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pHrResult {out} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.HResult); err != nil {
			return err
		}
	}
	// pulPercentCompleted {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.PercentCompleted); err != nil {
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

func (o *xxx_QueryStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pHrResult {out} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.HResult); err != nil {
			return err
		}
	}
	// pulPercentCompleted {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.PercentCompleted); err != nil {
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

// QueryStatusRequest structure represents the QueryStatus operation request
type QueryStatusRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryStatusRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryStatusOperation) *xxx_QueryStatusOperation {
	if op == nil {
		op = &xxx_QueryStatusOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *QueryStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryStatusOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryStatusResponse structure represents the QueryStatus operation response
type QueryStatusResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pHrResult: A pointer to a variable that receives the HRESULT that signals the current
	// state of the asynchronous operation.
	HResult int32 `idl:"name:pHrResult" json:"hresult"`
	// pulPercentCompleted: A pointer to a variable that receives the completion percentage
	// of the asynchronous operation. If the asynchronous operation is in progress, the
	// value MUST be between 0 and 99. If the operation has finished, the value MUST be
	// 100. If the progress of the operation cannot be estimated, the value MUST be 0.
	PercentCompleted uint32 `idl:"name:pulPercentCompleted" json:"percent_completed"`
	// Return: The QueryStatus return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryStatusResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryStatusOperation) *xxx_QueryStatusOperation {
	if op == nil {
		op = &xxx_QueryStatusOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.HResult = o.HResult
	op.PercentCompleted = o.PercentCompleted
	op.Return = o.Return
	return op
}

func (o *QueryStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryStatusOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.HResult = op.HResult
	o.PercentCompleted = op.PercentCompleted
	o.Return = op.Return
}
func (o *QueryStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
