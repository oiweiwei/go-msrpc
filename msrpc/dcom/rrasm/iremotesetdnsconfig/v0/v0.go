package iremotesetdnsconfig

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
	// IRemoteSetDnsConfig interface identifier 66a2db21-d706-11d0-a37b-00c04fc9da04
	RemoteSetDNSConfigIID = &dcom.IID{Data1: 0x66a2db21, Data2: 0xd706, Data3: 0x11d0, Data4: []byte{0xa3, 0x7b, 0x00, 0xc0, 0x4f, 0xc9, 0xda, 0x04}}
	// Syntax UUID
	RemoteSetDNSConfigSyntaxUUID = &uuid.UUID{TimeLow: 0x66a2db21, TimeMid: 0xd706, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x7b, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc9, 0xda, 0x4}}
	// Syntax ID
	RemoteSetDNSConfigSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RemoteSetDNSConfigSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRemoteSetDnsConfig interface.
type RemoteSetDNSConfigClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// SetDnsConfig operation.
	SetDNSConfig(context.Context, *SetDNSConfigRequest, ...dcerpc.CallOption) (*SetDNSConfigResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RemoteSetDNSConfigClient
}

type xxx_DefaultRemoteSetDNSConfigClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRemoteSetDNSConfigClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultRemoteSetDNSConfigClient) SetDNSConfig(ctx context.Context, in *SetDNSConfigRequest, opts ...dcerpc.CallOption) (*SetDNSConfigResponse, error) {
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
	out := &SetDNSConfigResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteSetDNSConfigClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteSetDNSConfigClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRemoteSetDNSConfigClient) IPID(ctx context.Context, ipid *dcom.IPID) RemoteSetDNSConfigClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRemoteSetDNSConfigClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewRemoteSetDNSConfigClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteSetDNSConfigClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteSetDNSConfigSyntaxV0_0))...)
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
	return &xxx_DefaultRemoteSetDNSConfigClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_SetDNSConfigOperation structure represents the SetDnsConfig operation
type xxx_SetDNSConfigOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConfigID uint32         `idl:"name:dwConfigId" json:"config_id"`
	NewValue uint32         `idl:"name:dwNewValue" json:"new_value"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDNSConfigOperation) OpNum() int { return 3 }

func (o *xxx_SetDNSConfigOperation) OpName() string { return "/IRemoteSetDnsConfig/v0/SetDnsConfig" }

func (o *xxx_SetDNSConfigOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDNSConfigOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwConfigId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ConfigID); err != nil {
			return err
		}
	}
	// dwNewValue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NewValue); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDNSConfigOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwConfigId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ConfigID); err != nil {
			return err
		}
	}
	// dwNewValue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NewValue); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDNSConfigOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDNSConfigOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetDNSConfigOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetDNSConfigRequest structure represents the SetDnsConfig operation request
type SetDNSConfigRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	ConfigID uint32         `idl:"name:dwConfigId" json:"config_id"`
	NewValue uint32         `idl:"name:dwNewValue" json:"new_value"`
}

func (o *SetDNSConfigRequest) xxx_ToOp(ctx context.Context, op *xxx_SetDNSConfigOperation) *xxx_SetDNSConfigOperation {
	if op == nil {
		op = &xxx_SetDNSConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ConfigID = o.ConfigID
	op.NewValue = o.NewValue
	return op
}

func (o *SetDNSConfigRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDNSConfigOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConfigID = op.ConfigID
	o.NewValue = op.NewValue
}
func (o *SetDNSConfigRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetDNSConfigRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDNSConfigOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDNSConfigResponse structure represents the SetDnsConfig operation response
type SetDNSConfigResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetDnsConfig return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDNSConfigResponse) xxx_ToOp(ctx context.Context, op *xxx_SetDNSConfigOperation) *xxx_SetDNSConfigOperation {
	if op == nil {
		op = &xxx_SetDNSConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetDNSConfigResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDNSConfigOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDNSConfigResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetDNSConfigResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDNSConfigOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
