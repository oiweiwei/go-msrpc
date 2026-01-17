package iremoteicficsconfig

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rrasm"
)

var (
	// IRemoteICFICSConfig interface identifier 66a2db22-d706-11d0-a37b-00c04fc9da04
	RemoteICFICSConfigIID = &dcom.IID{Data1: 0x66a2db22, Data2: 0xd706, Data3: 0x11d0, Data4: []byte{0xa3, 0x7b, 0x00, 0xc0, 0x4f, 0xc9, 0xda, 0x04}}
	// Syntax UUID
	RemoteICFICSConfigSyntaxUUID = &uuid.UUID{TimeLow: 0x66a2db22, TimeMid: 0xd706, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x7b, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc9, 0xda, 0x4}}
	// Syntax ID
	RemoteICFICSConfigSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RemoteICFICSConfigSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRemoteICFICSConfig interface.
type RemoteICFICSConfigClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetIcfEnabled operation.
	GetICFEnabled(context.Context, *GetICFEnabledRequest, ...dcerpc.CallOption) (*GetICFEnabledResponse, error)

	// GetIcsEnabled operation.
	GetICSEnabled(context.Context, *GetICSEnabledRequest, ...dcerpc.CallOption) (*GetICSEnabledResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RemoteICFICSConfigClient
}

type xxx_DefaultRemoteICFICSConfigClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRemoteICFICSConfigClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultRemoteICFICSConfigClient) GetICFEnabled(ctx context.Context, in *GetICFEnabledRequest, opts ...dcerpc.CallOption) (*GetICFEnabledResponse, error) {
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
	out := &GetICFEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteICFICSConfigClient) GetICSEnabled(ctx context.Context, in *GetICSEnabledRequest, opts ...dcerpc.CallOption) (*GetICSEnabledResponse, error) {
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
	out := &GetICSEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteICFICSConfigClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteICFICSConfigClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRemoteICFICSConfigClient) IPID(ctx context.Context, ipid *dcom.IPID) RemoteICFICSConfigClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRemoteICFICSConfigClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewRemoteICFICSConfigClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteICFICSConfigClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteICFICSConfigSyntaxV0_0))...)
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
	return &xxx_DefaultRemoteICFICSConfigClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetICFEnabledOperation structure represents the GetIcfEnabled operation
type xxx_GetICFEnabledOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Status bool           `idl:"name:status" json:"status"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetICFEnabledOperation) OpNum() int { return 3 }

func (o *xxx_GetICFEnabledOperation) OpName() string { return "/IRemoteICFICSConfig/v0/GetIcfEnabled" }

func (o *xxx_GetICFEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetICFEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetICFEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetICFEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetICFEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetICFEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetICFEnabledRequest structure represents the GetIcfEnabled operation request
type GetICFEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetICFEnabledRequest) xxx_ToOp(ctx context.Context, op *xxx_GetICFEnabledOperation) *xxx_GetICFEnabledOperation {
	if op == nil {
		op = &xxx_GetICFEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetICFEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_GetICFEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetICFEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetICFEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetICFEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetICFEnabledResponse structure represents the GetIcfEnabled operation response
type GetICFEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Status bool           `idl:"name:status" json:"status"`
	// Return: The GetIcfEnabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetICFEnabledResponse) xxx_ToOp(ctx context.Context, op *xxx_GetICFEnabledOperation) *xxx_GetICFEnabledOperation {
	if op == nil {
		op = &xxx_GetICFEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Status = o.Status
	op.Return = o.Return
	return op
}

func (o *GetICFEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_GetICFEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Status = op.Status
	o.Return = op.Return
}
func (o *GetICFEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetICFEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetICFEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetICSEnabledOperation structure represents the GetIcsEnabled operation
type xxx_GetICSEnabledOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Status bool           `idl:"name:status" json:"status"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetICSEnabledOperation) OpNum() int { return 4 }

func (o *xxx_GetICSEnabledOperation) OpName() string { return "/IRemoteICFICSConfig/v0/GetIcsEnabled" }

func (o *xxx_GetICSEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetICSEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetICSEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetICSEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetICSEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetICSEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetICSEnabledRequest structure represents the GetIcsEnabled operation request
type GetICSEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetICSEnabledRequest) xxx_ToOp(ctx context.Context, op *xxx_GetICSEnabledOperation) *xxx_GetICSEnabledOperation {
	if op == nil {
		op = &xxx_GetICSEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetICSEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_GetICSEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetICSEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetICSEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetICSEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetICSEnabledResponse structure represents the GetIcsEnabled operation response
type GetICSEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Status bool           `idl:"name:status" json:"status"`
	// Return: The GetIcsEnabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetICSEnabledResponse) xxx_ToOp(ctx context.Context, op *xxx_GetICSEnabledOperation) *xxx_GetICSEnabledOperation {
	if op == nil {
		op = &xxx_GetICSEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Status = o.Status
	op.Return = o.Return
	return op
}

func (o *GetICSEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_GetICSEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Status = op.Status
	o.Return = op.Return
}
func (o *GetICSEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetICSEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetICSEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
