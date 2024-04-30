package icapabilitysupport

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	coma "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = coma.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

var (
	// ICapabilitySupport interface identifier 47cde9a1-0bf6-11d2-8016-00c04fb9988e
	CapabilitySupportIID = &dcom.IID{Data1: 0x47cde9a1, Data2: 0x0bf6, Data3: 0x11d2, Data4: []byte{0x80, 0x16, 0x00, 0xc0, 0x4f, 0xb9, 0x98, 0x8e}}
	// Syntax UUID
	CapabilitySupportSyntaxUUID = &uuid.UUID{TimeLow: 0x47cde9a1, TimeMid: 0xbf6, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0x16, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb9, 0x98, 0x8e}}
	// Syntax ID
	CapabilitySupportSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CapabilitySupportSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICapabilitySupport interface.
type CapabilitySupportClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to start instance load balancing.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	Start(context.Context, *StartRequest, ...dcerpc.CallOption) (*StartResponse, error)

	// This method is called by a client to stop instance load balancing.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	Stop(context.Context, *StopRequest, ...dcerpc.CallOption) (*StopResponse, error)

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// This method is called by a client to determine whether instance load balancing is
	// installed.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	IsInstalled(context.Context, *IsInstalledRequest, ...dcerpc.CallOption) (*IsInstalledResponse, error)

	// This method is called by a client to determine whether instance load balancing is
	// running.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	IsRunning(context.Context, *IsRunningRequest, ...dcerpc.CallOption) (*IsRunningResponse, error)

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CapabilitySupportClient
}

type xxx_DefaultCapabilitySupportClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCapabilitySupportClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultCapabilitySupportClient) Start(ctx context.Context, in *StartRequest, opts ...dcerpc.CallOption) (*StartResponse, error) {
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
	out := &StartResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCapabilitySupportClient) Stop(ctx context.Context, in *StopRequest, opts ...dcerpc.CallOption) (*StopResponse, error) {
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
	out := &StopResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCapabilitySupportClient) IsInstalled(ctx context.Context, in *IsInstalledRequest, opts ...dcerpc.CallOption) (*IsInstalledResponse, error) {
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
	out := &IsInstalledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCapabilitySupportClient) IsRunning(ctx context.Context, in *IsRunningRequest, opts ...dcerpc.CallOption) (*IsRunningResponse, error) {
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
	out := &IsRunningResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCapabilitySupportClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCapabilitySupportClient) IPID(ctx context.Context, ipid *dcom.IPID) CapabilitySupportClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCapabilitySupportClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewCapabilitySupportClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CapabilitySupportClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CapabilitySupportSyntaxV0_0))...)
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
	return &xxx_DefaultCapabilitySupportClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_StartOperation structure represents the Start operation
type xxx_StartOperation struct {
	This   *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat         `idl:"name:That" json:"that"`
	CSS    coma.CatServerServices `idl:"name:i_css" json:"css"`
	Return int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_StartOperation) OpNum() int { return 3 }

func (o *xxx_StartOperation) OpName() string { return "/ICapabilitySupport/v0/Start" }

func (o *xxx_StartOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// i_css {in} (1:{alias=CatSrvServices}(enum))
	{
		if err := w.WriteData(uint16(o.CSS)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// i_css {in} (1:{alias=CatSrvServices}(enum))
	{
		if err := w.ReadData((*uint16)(&o.CSS)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_StartOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// StartRequest structure represents the Start operation request
type StartRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// i_css: MUST be set to css_lb (see section 2.2.7).
	CSS coma.CatServerServices `idl:"name:i_css" json:"css"`
}

func (o *StartRequest) xxx_ToOp(ctx context.Context) *xxx_StartOperation {
	if o == nil {
		return &xxx_StartOperation{}
	}
	return &xxx_StartOperation{
		This: o.This,
		CSS:  o.CSS,
	}
}

func (o *StartRequest) xxx_FromOp(ctx context.Context, op *xxx_StartOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CSS = op.CSS
}
func (o *StartRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *StartRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StartResponse structure represents the Start operation response
type StartResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Start return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StartResponse) xxx_ToOp(ctx context.Context) *xxx_StartOperation {
	if o == nil {
		return &xxx_StartOperation{}
	}
	return &xxx_StartOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *StartResponse) xxx_FromOp(ctx context.Context, op *xxx_StartOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *StartResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *StartResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StopOperation structure represents the Stop operation
type xxx_StopOperation struct {
	This   *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat         `idl:"name:That" json:"that"`
	CSS    coma.CatServerServices `idl:"name:i_css" json:"css"`
	Return int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_StopOperation) OpNum() int { return 4 }

func (o *xxx_StopOperation) OpName() string { return "/ICapabilitySupport/v0/Stop" }

func (o *xxx_StopOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// i_css {in} (1:{alias=CatSrvServices}(enum))
	{
		if err := w.WriteData(uint16(o.CSS)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// i_css {in} (1:{alias=CatSrvServices}(enum))
	{
		if err := w.ReadData((*uint16)(&o.CSS)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_StopOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// StopRequest structure represents the Stop operation request
type StopRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// i_css: MUST be set to css_lb (see section 2.2.7).
	CSS coma.CatServerServices `idl:"name:i_css" json:"css"`
}

func (o *StopRequest) xxx_ToOp(ctx context.Context) *xxx_StopOperation {
	if o == nil {
		return &xxx_StopOperation{}
	}
	return &xxx_StopOperation{
		This: o.This,
		CSS:  o.CSS,
	}
}

func (o *StopRequest) xxx_FromOp(ctx context.Context, op *xxx_StopOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CSS = op.CSS
}
func (o *StopRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *StopRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StopResponse structure represents the Stop operation response
type StopResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Stop return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StopResponse) xxx_ToOp(ctx context.Context) *xxx_StopOperation {
	if o == nil {
		return &xxx_StopOperation{}
	}
	return &xxx_StopOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *StopResponse) xxx_FromOp(ctx context.Context, op *xxx_StopOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *StopResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *StopResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsInstalledOperation structure represents the IsInstalled operation
type xxx_IsInstalledOperation struct {
	This   *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat         `idl:"name:That" json:"that"`
	CSS    coma.CatServerServices `idl:"name:i_css" json:"css"`
	Status uint32                 `idl:"name:pulStatus" json:"status"`
	Return int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_IsInstalledOperation) OpNum() int { return 7 }

func (o *xxx_IsInstalledOperation) OpName() string { return "/ICapabilitySupport/v0/IsInstalled" }

func (o *xxx_IsInstalledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsInstalledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// i_css {in} (1:{alias=CatSrvServices}(enum))
	{
		if err := w.WriteData(uint16(o.CSS)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsInstalledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// i_css {in} (1:{alias=CatSrvServices}(enum))
	{
		if err := w.ReadData((*uint16)(&o.CSS)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsInstalledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsInstalledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pulStatus {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Status); err != nil {
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

func (o *xxx_IsInstalledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pulStatus {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Status); err != nil {
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

// IsInstalledRequest structure represents the IsInstalled operation request
type IsInstalledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// i_css: MUST be set to css_lb (see section 2.2.7).
	CSS coma.CatServerServices `idl:"name:i_css" json:"css"`
}

func (o *IsInstalledRequest) xxx_ToOp(ctx context.Context) *xxx_IsInstalledOperation {
	if o == nil {
		return &xxx_IsInstalledOperation{}
	}
	return &xxx_IsInstalledOperation{
		This: o.This,
		CSS:  o.CSS,
	}
}

func (o *IsInstalledRequest) xxx_FromOp(ctx context.Context, op *xxx_IsInstalledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CSS = op.CSS
}
func (o *IsInstalledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsInstalledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsInstalledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsInstalledResponse structure represents the IsInstalled operation response
type IsInstalledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pulStatus: A pointer to a variable that, upon a successful return, MUST be set to
	// TRUE (0x00000001) or FALSE (0x000000000) to indicate whether component load balancing
	// support is installed.
	Status uint32 `idl:"name:pulStatus" json:"status"`
	// Return: The IsInstalled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsInstalledResponse) xxx_ToOp(ctx context.Context) *xxx_IsInstalledOperation {
	if o == nil {
		return &xxx_IsInstalledOperation{}
	}
	return &xxx_IsInstalledOperation{
		That:   o.That,
		Status: o.Status,
		Return: o.Return,
	}
}

func (o *IsInstalledResponse) xxx_FromOp(ctx context.Context, op *xxx_IsInstalledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Status = op.Status
	o.Return = op.Return
}
func (o *IsInstalledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsInstalledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsInstalledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsRunningOperation structure represents the IsRunning operation
type xxx_IsRunningOperation struct {
	This   *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat             `idl:"name:That" json:"that"`
	CSS    coma.CatServerServices     `idl:"name:i_css" json:"css"`
	States coma.CatServerServiceState `idl:"name:pulStates" json:"states"`
	Return int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_IsRunningOperation) OpNum() int { return 8 }

func (o *xxx_IsRunningOperation) OpName() string { return "/ICapabilitySupport/v0/IsRunning" }

func (o *xxx_IsRunningOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsRunningOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// i_css {in} (1:{alias=CatSrvServices}(enum))
	{
		if err := w.WriteData(uint16(o.CSS)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsRunningOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// i_css {in} (1:{alias=CatSrvServices}(enum))
	{
		if err := w.ReadData((*uint16)(&o.CSS)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsRunningOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsRunningOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pulStates {out} (1:{pointer=ref}*(1))(2:{alias=CatSrvServiceState}(enum))
	{
		if err := w.WriteData(uint16(o.States)); err != nil {
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

func (o *xxx_IsRunningOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pulStates {out} (1:{pointer=ref}*(1))(2:{alias=CatSrvServiceState}(enum))
	{
		if err := w.ReadData((*uint16)(&o.States)); err != nil {
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

// IsRunningRequest structure represents the IsRunning operation request
type IsRunningRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// i_css: MUST be set to css_lb (see section 2.2.7).
	CSS coma.CatServerServices `idl:"name:i_css" json:"css"`
}

func (o *IsRunningRequest) xxx_ToOp(ctx context.Context) *xxx_IsRunningOperation {
	if o == nil {
		return &xxx_IsRunningOperation{}
	}
	return &xxx_IsRunningOperation{
		This: o.This,
		CSS:  o.CSS,
	}
}

func (o *IsRunningRequest) xxx_FromOp(ctx context.Context, op *xxx_IsRunningOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CSS = op.CSS
}
func (o *IsRunningRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsRunningRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsRunningOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsRunningResponse structure represents the IsRunning operation response
type IsRunningResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pulStates: A pointer to a variable that, upon a successful return, MUST be set to
	// the CatSrvServiceState (section 2.2.8) value that indicates the current running state
	// of instance load balancing.
	States coma.CatServerServiceState `idl:"name:pulStates" json:"states"`
	// Return: The IsRunning return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsRunningResponse) xxx_ToOp(ctx context.Context) *xxx_IsRunningOperation {
	if o == nil {
		return &xxx_IsRunningOperation{}
	}
	return &xxx_IsRunningOperation{
		That:   o.That,
		States: o.States,
		Return: o.Return,
	}
}

func (o *IsRunningResponse) xxx_FromOp(ctx context.Context, op *xxx_IsRunningOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.States = op.States
	o.Return = op.Return
}
func (o *IsRunningResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsRunningResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsRunningOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
