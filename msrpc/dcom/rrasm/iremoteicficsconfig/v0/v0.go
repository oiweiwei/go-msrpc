package iremoteicficsconfig

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
)

var (
	// import guard
	GoPackage = "dcom/rrasm"
)

var (
	// IRemoteICFICSConfig interface identifier 66a2db22-d706-11d0-a37b-00c04fc9da04
	RemoteIcficsConfigIID = &dcom.IID{Data1: 0x66a2db22, Data2: 0xd706, Data3: 0x11d0, Data4: []byte{0xa3, 0x7b, 0x00, 0xc0, 0x4f, 0xc9, 0xda, 0x04}}
	// Syntax UUID
	RemoteIcficsConfigSyntaxUUID = &uuid.UUID{TimeLow: 0x66a2db22, TimeMid: 0xd706, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x7b, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc9, 0xda, 0x4}}
	// Syntax ID
	RemoteIcficsConfigSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RemoteIcficsConfigSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRemoteICFICSConfig interface.
type RemoteIcficsConfigClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetIcfEnabled operation.
	GetIcfEnabled(context.Context, *GetIcfEnabledRequest, ...dcerpc.CallOption) (*GetIcfEnabledResponse, error)

	// GetIcsEnabled operation.
	GetIcsEnabled(context.Context, *GetIcsEnabledRequest, ...dcerpc.CallOption) (*GetIcsEnabledResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RemoteIcficsConfigClient
}

type xxx_DefaultRemoteIcficsConfigClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRemoteIcficsConfigClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultRemoteIcficsConfigClient) GetIcfEnabled(ctx context.Context, in *GetIcfEnabledRequest, opts ...dcerpc.CallOption) (*GetIcfEnabledResponse, error) {
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
	out := &GetIcfEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteIcficsConfigClient) GetIcsEnabled(ctx context.Context, in *GetIcsEnabledRequest, opts ...dcerpc.CallOption) (*GetIcsEnabledResponse, error) {
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
	out := &GetIcsEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteIcficsConfigClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteIcficsConfigClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRemoteIcficsConfigClient) IPID(ctx context.Context, ipid *dcom.IPID) RemoteIcficsConfigClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRemoteIcficsConfigClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewRemoteIcficsConfigClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteIcficsConfigClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteIcficsConfigSyntaxV0_0))...)
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
	return &xxx_DefaultRemoteIcficsConfigClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetIcfEnabledOperation structure represents the GetIcfEnabled operation
type xxx_GetIcfEnabledOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Status bool           `idl:"name:status" json:"status"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIcfEnabledOperation) OpNum() int { return 3 }

func (o *xxx_GetIcfEnabledOperation) OpName() string { return "/IRemoteICFICSConfig/v0/GetIcfEnabled" }

func (o *xxx_GetIcfEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIcfEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIcfEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIcfEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIcfEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// status {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.Status {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
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

func (o *xxx_GetIcfEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// status {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bStatus int32
		if err := w.ReadData(&_bStatus); err != nil {
			return err
		}
		o.Status = _bStatus != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetIcfEnabledRequest structure represents the GetIcfEnabled operation request
type GetIcfEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIcfEnabledRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIcfEnabledOperation) *xxx_GetIcfEnabledOperation {
	if op == nil {
		op = &xxx_GetIcfEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIcfEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIcfEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIcfEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIcfEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIcfEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIcfEnabledResponse structure represents the GetIcfEnabled operation response
type GetIcfEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Status bool           `idl:"name:status" json:"status"`
	// Return: The GetIcfEnabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIcfEnabledResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIcfEnabledOperation) *xxx_GetIcfEnabledOperation {
	if op == nil {
		op = &xxx_GetIcfEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Status = o.Status
	op.Return = o.Return
	return op
}

func (o *GetIcfEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIcfEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Status = op.Status
	o.Return = op.Return
}
func (o *GetIcfEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIcfEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIcfEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIcsEnabledOperation structure represents the GetIcsEnabled operation
type xxx_GetIcsEnabledOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Status bool           `idl:"name:status" json:"status"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIcsEnabledOperation) OpNum() int { return 4 }

func (o *xxx_GetIcsEnabledOperation) OpName() string { return "/IRemoteICFICSConfig/v0/GetIcsEnabled" }

func (o *xxx_GetIcsEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIcsEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIcsEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIcsEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIcsEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// status {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.Status {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
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

func (o *xxx_GetIcsEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// status {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bStatus int32
		if err := w.ReadData(&_bStatus); err != nil {
			return err
		}
		o.Status = _bStatus != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetIcsEnabledRequest structure represents the GetIcsEnabled operation request
type GetIcsEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIcsEnabledRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIcsEnabledOperation) *xxx_GetIcsEnabledOperation {
	if op == nil {
		op = &xxx_GetIcsEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIcsEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIcsEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIcsEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIcsEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIcsEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIcsEnabledResponse structure represents the GetIcsEnabled operation response
type GetIcsEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Status bool           `idl:"name:status" json:"status"`
	// Return: The GetIcsEnabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIcsEnabledResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIcsEnabledOperation) *xxx_GetIcsEnabledOperation {
	if op == nil {
		op = &xxx_GetIcsEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Status = o.Status
	op.Return = o.Return
	return op
}

func (o *GetIcsEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIcsEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Status = op.Status
	o.Return = op.Return
}
func (o *GetIcsEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIcsEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIcsEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
