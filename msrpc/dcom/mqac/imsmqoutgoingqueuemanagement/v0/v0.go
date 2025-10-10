package imsmqoutgoingqueuemanagement

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	mqac "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac"
	imsmqmanagement "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqmanagement/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = imsmqmanagement.GoPackage
	_ = oaut.GoPackage
	_ = mqac.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQOutgoingQueueManagement interface identifier 64c478fb-f9b0-4695-8a7f-439ac94326d3
	ImsmqOutgoingQueueManagementIID = &dcom.IID{Data1: 0x64c478fb, Data2: 0xf9b0, Data3: 0x4695, Data4: []byte{0x8a, 0x7f, 0x43, 0x9a, 0xc9, 0x43, 0x26, 0xd3}}
	// Syntax UUID
	ImsmqOutgoingQueueManagementSyntaxUUID = &uuid.UUID{TimeLow: 0x64c478fb, TimeMid: 0xf9b0, TimeHiAndVersion: 0x4695, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x7f, Node: [6]uint8{0x43, 0x9a, 0xc9, 0x43, 0x26, 0xd3}}
	// Syntax ID
	ImsmqOutgoingQueueManagementSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ImsmqOutgoingQueueManagementSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQOutgoingQueueManagement interface.
type ImsmqOutgoingQueueManagementClient interface {

	// IMSMQManagement retrieval method.
	ImsmqManagement() imsmqmanagement.ImsmqManagementClient

	// State operation.
	GetState(context.Context, *GetStateRequest, ...dcerpc.CallOption) (*GetStateResponse, error)

	// NextHops operation.
	GetNextHops(context.Context, *GetNextHopsRequest, ...dcerpc.CallOption) (*GetNextHopsResponse, error)

	// EodGetSendInfo operation.
	EodGetSendInfo(context.Context, *EodGetSendInfoRequest, ...dcerpc.CallOption) (*EodGetSendInfoResponse, error)

	// Resume operation.
	Resume(context.Context, *ResumeRequest, ...dcerpc.CallOption) (*ResumeResponse, error)

	// Pause operation.
	Pause(context.Context, *PauseRequest, ...dcerpc.CallOption) (*PauseResponse, error)

	// EodResend operation.
	EodResend(context.Context, *EodResendRequest, ...dcerpc.CallOption) (*EodResendResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ImsmqOutgoingQueueManagementClient
}

type xxx_DefaultImsmqOutgoingQueueManagementClient struct {
	imsmqmanagement.ImsmqManagementClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultImsmqOutgoingQueueManagementClient) ImsmqManagement() imsmqmanagement.ImsmqManagementClient {
	return o.ImsmqManagementClient
}

func (o *xxx_DefaultImsmqOutgoingQueueManagementClient) GetState(ctx context.Context, in *GetStateRequest, opts ...dcerpc.CallOption) (*GetStateResponse, error) {
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
	out := &GetStateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqOutgoingQueueManagementClient) GetNextHops(ctx context.Context, in *GetNextHopsRequest, opts ...dcerpc.CallOption) (*GetNextHopsResponse, error) {
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
	out := &GetNextHopsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqOutgoingQueueManagementClient) EodGetSendInfo(ctx context.Context, in *EodGetSendInfoRequest, opts ...dcerpc.CallOption) (*EodGetSendInfoResponse, error) {
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
	out := &EodGetSendInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqOutgoingQueueManagementClient) Resume(ctx context.Context, in *ResumeRequest, opts ...dcerpc.CallOption) (*ResumeResponse, error) {
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
	out := &ResumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqOutgoingQueueManagementClient) Pause(ctx context.Context, in *PauseRequest, opts ...dcerpc.CallOption) (*PauseResponse, error) {
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
	out := &PauseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqOutgoingQueueManagementClient) EodResend(ctx context.Context, in *EodResendRequest, opts ...dcerpc.CallOption) (*EodResendResponse, error) {
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
	out := &EodResendResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqOutgoingQueueManagementClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultImsmqOutgoingQueueManagementClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultImsmqOutgoingQueueManagementClient) IPID(ctx context.Context, ipid *dcom.IPID) ImsmqOutgoingQueueManagementClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultImsmqOutgoingQueueManagementClient{
		ImsmqManagementClient: o.ImsmqManagementClient.IPID(ctx, ipid),
		cc:                    o.cc,
		ipid:                  ipid,
	}
}

func NewImsmqOutgoingQueueManagementClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ImsmqOutgoingQueueManagementClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ImsmqOutgoingQueueManagementSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := imsmqmanagement.NewImsmqManagementClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultImsmqOutgoingQueueManagementClient{
		ImsmqManagementClient: base,
		cc:                    cc,
		ipid:                  ipid,
	}, nil
}

// xxx_GetStateOperation structure represents the State operation
type xxx_GetStateOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	State  int32          `idl:"name:plState" json:"state"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetStateOperation) OpNum() int { return 16 }

func (o *xxx_GetStateOperation) OpName() string { return "/IMSMQOutgoingQueueManagement/v0/State" }

func (o *xxx_GetStateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetStateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetStateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plState {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.State); err != nil {
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

func (o *xxx_GetStateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plState {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.State); err != nil {
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

// GetStateRequest structure represents the State operation request
type GetStateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetStateRequest) xxx_ToOp(ctx context.Context, op *xxx_GetStateOperation) *xxx_GetStateOperation {
	if op == nil {
		op = &xxx_GetStateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetStateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetStateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetStateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetStateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetStateResponse structure represents the State operation response
type GetStateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	State int32          `idl:"name:plState" json:"state"`
	// Return: The State return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetStateResponse) xxx_ToOp(ctx context.Context, op *xxx_GetStateOperation) *xxx_GetStateOperation {
	if op == nil {
		op = &xxx_GetStateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.State = o.State
	op.Return = o.Return
	return op
}

func (o *GetStateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetStateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.State = op.State
	o.Return = op.Return
}
func (o *GetStateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetStateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNextHopsOperation structure represents the NextHops operation
type xxx_GetNextHopsOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	NextHops *oaut.Variant  `idl:"name:pvNextHops" json:"next_hops"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNextHopsOperation) OpNum() int { return 17 }

func (o *xxx_GetNextHopsOperation) OpName() string {
	return "/IMSMQOutgoingQueueManagement/v0/NextHops"
}

func (o *xxx_GetNextHopsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextHopsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetNextHopsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetNextHopsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextHopsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvNextHops {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.NextHops != nil {
			_ptr_pvNextHops := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NextHops != nil {
					if err := o.NextHops.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NextHops, _ptr_pvNextHops); err != nil {
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

func (o *xxx_GetNextHopsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvNextHops {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvNextHops := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NextHops == nil {
				o.NextHops = &oaut.Variant{}
			}
			if err := o.NextHops.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvNextHops := func(ptr interface{}) { o.NextHops = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.NextHops, _s_pvNextHops, _ptr_pvNextHops); err != nil {
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

// GetNextHopsRequest structure represents the NextHops operation request
type GetNextHopsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetNextHopsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNextHopsOperation) *xxx_GetNextHopsOperation {
	if op == nil {
		op = &xxx_GetNextHopsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetNextHopsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNextHopsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetNextHopsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNextHopsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNextHopsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNextHopsResponse structure represents the NextHops operation response
type GetNextHopsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	NextHops *oaut.Variant  `idl:"name:pvNextHops" json:"next_hops"`
	// Return: The NextHops return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNextHopsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNextHopsOperation) *xxx_GetNextHopsOperation {
	if op == nil {
		op = &xxx_GetNextHopsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.NextHops = o.NextHops
	op.Return = o.Return
	return op
}

func (o *GetNextHopsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNextHopsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.NextHops = op.NextHops
	o.Return = op.Return
}
func (o *GetNextHopsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNextHopsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNextHopsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EodGetSendInfoOperation structure represents the EodGetSendInfo operation
type xxx_EodGetSendInfoOperation struct {
	This       *dcom.ORPCThis        `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat        `idl:"name:That" json:"that"`
	Collection *mqac.ImsmqCollection `idl:"name:ppCollection" json:"collection"`
	Return     int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_EodGetSendInfoOperation) OpNum() int { return 18 }

func (o *xxx_EodGetSendInfoOperation) OpName() string {
	return "/IMSMQOutgoingQueueManagement/v0/EodGetSendInfo"
}

func (o *xxx_EodGetSendInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EodGetSendInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EodGetSendInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_EodGetSendInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EodGetSendInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppCollection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQCollection}(interface))
	{
		if o.Collection != nil {
			_ptr_ppCollection := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Collection != nil {
					if err := o.Collection.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.ImsmqCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Collection, _ptr_ppCollection); err != nil {
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

func (o *xxx_EodGetSendInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppCollection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQCollection}(interface))
	{
		_ptr_ppCollection := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Collection == nil {
				o.Collection = &mqac.ImsmqCollection{}
			}
			if err := o.Collection.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppCollection := func(ptr interface{}) { o.Collection = *ptr.(**mqac.ImsmqCollection) }
		if err := w.ReadPointer(&o.Collection, _s_ppCollection, _ptr_ppCollection); err != nil {
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

// EodGetSendInfoRequest structure represents the EodGetSendInfo operation request
type EodGetSendInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *EodGetSendInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_EodGetSendInfoOperation) *xxx_EodGetSendInfoOperation {
	if op == nil {
		op = &xxx_EodGetSendInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *EodGetSendInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_EodGetSendInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *EodGetSendInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EodGetSendInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EodGetSendInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EodGetSendInfoResponse structure represents the EodGetSendInfo operation response
type EodGetSendInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat        `idl:"name:That" json:"that"`
	Collection *mqac.ImsmqCollection `idl:"name:ppCollection" json:"collection"`
	// Return: The EodGetSendInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EodGetSendInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_EodGetSendInfoOperation) *xxx_EodGetSendInfoOperation {
	if op == nil {
		op = &xxx_EodGetSendInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Collection = o.Collection
	op.Return = o.Return
	return op
}

func (o *EodGetSendInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_EodGetSendInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Collection = op.Collection
	o.Return = op.Return
}
func (o *EodGetSendInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EodGetSendInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EodGetSendInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResumeOperation structure represents the Resume operation
type xxx_ResumeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ResumeOperation) OpNum() int { return 19 }

func (o *xxx_ResumeOperation) OpName() string { return "/IMSMQOutgoingQueueManagement/v0/Resume" }

func (o *xxx_ResumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ResumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ResumeRequest structure represents the Resume operation request
type ResumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ResumeRequest) xxx_ToOp(ctx context.Context, op *xxx_ResumeOperation) *xxx_ResumeOperation {
	if op == nil {
		op = &xxx_ResumeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ResumeRequest) xxx_FromOp(ctx context.Context, op *xxx_ResumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ResumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ResumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResumeResponse structure represents the Resume operation response
type ResumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Resume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ResumeResponse) xxx_ToOp(ctx context.Context, op *xxx_ResumeOperation) *xxx_ResumeOperation {
	if op == nil {
		op = &xxx_ResumeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ResumeResponse) xxx_FromOp(ctx context.Context, op *xxx_ResumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ResumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ResumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PauseOperation structure represents the Pause operation
type xxx_PauseOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PauseOperation) OpNum() int { return 20 }

func (o *xxx_PauseOperation) OpName() string { return "/IMSMQOutgoingQueueManagement/v0/Pause" }

func (o *xxx_PauseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PauseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_PauseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PauseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// PauseRequest structure represents the Pause operation request
type PauseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *PauseRequest) xxx_ToOp(ctx context.Context, op *xxx_PauseOperation) *xxx_PauseOperation {
	if op == nil {
		op = &xxx_PauseOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *PauseRequest) xxx_FromOp(ctx context.Context, op *xxx_PauseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *PauseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PauseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PauseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PauseResponse structure represents the Pause operation response
type PauseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Pause return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PauseResponse) xxx_ToOp(ctx context.Context, op *xxx_PauseOperation) *xxx_PauseOperation {
	if op == nil {
		op = &xxx_PauseOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *PauseResponse) xxx_FromOp(ctx context.Context, op *xxx_PauseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PauseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PauseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PauseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EodResendOperation structure represents the EodResend operation
type xxx_EodResendOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EodResendOperation) OpNum() int { return 21 }

func (o *xxx_EodResendOperation) OpName() string { return "/IMSMQOutgoingQueueManagement/v0/EodResend" }

func (o *xxx_EodResendOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EodResendOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EodResendOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_EodResendOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EodResendOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EodResendOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// EodResendRequest structure represents the EodResend operation request
type EodResendRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *EodResendRequest) xxx_ToOp(ctx context.Context, op *xxx_EodResendOperation) *xxx_EodResendOperation {
	if op == nil {
		op = &xxx_EodResendOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *EodResendRequest) xxx_FromOp(ctx context.Context, op *xxx_EodResendOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *EodResendRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EodResendRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EodResendOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EodResendResponse structure represents the EodResend operation response
type EodResendResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EodResend return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EodResendResponse) xxx_ToOp(ctx context.Context, op *xxx_EodResendOperation) *xxx_EodResendOperation {
	if op == nil {
		op = &xxx_EodResendOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *EodResendResponse) xxx_FromOp(ctx context.Context, op *xxx_EodResendOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *EodResendResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EodResendResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EodResendOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
